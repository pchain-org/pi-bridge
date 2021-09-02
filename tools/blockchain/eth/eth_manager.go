package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	cctypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	ctypes "github.com/cosmos/cosmos-sdk/types"
	typestx "github.com/cosmos/cosmos-sdk/types/tx"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/config"
	"github.com/pchain-org/pi-bridge/tools/log"
	"github.com/pchain-org/pi-bridge/tools/types"
	blockmoduletypes "github.com/pchain-org/pi-bridge/x/block/types"
	trxtypes "github.com/pchain-org/pi-bridge/x/trx/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	tdmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
	tdmtypes "github.com/tendermint/tendermint/types"
	pbridge_abi "github.com/terry108/cross-chain-contract/pbridge_abi"
	"google.golang.org/grpc"
)

const (
	ChanLen = 64
)

type EthManager struct {
	cfg          *config.ServiceConfig
	restClient   *common.RestClient
	ethClient    *ethclient.Client
	bridgeClient *http.HTTP
	// currentHeight uint64
	// forceHeight   uint64

	CMPrivk  cctypes.PrivKey
	CMAcc    ctypes.AccAddress
	CMSeq    *CosmosSeq
	CMAccNum uint64
	CMFees   ctypes.Coins
	CMGas    uint64
	CMCdc    *codec.AminoCodec

	senders []*EthSender
	cdc     codec.ProtoCodecMarshaler
}
type CosmosSeq struct {
	lock sync.Mutex
	val  uint64
}

func (seq *CosmosSeq) GetAndAdd() uint64 {
	seq.lock.Lock()
	defer func() {
		seq.val += 1
		seq.lock.Unlock()
	}()
	return seq.val
}

func NewEthManager(ctx *ctypes.Context, cdc codec.ProtoCodecMarshaler, cfg *config.ServiceConfig, authKeeper *authkeeper.AccountKeeper) (*EthManager, error) {
	restClient := common.NewRestClient()
	ethClient, err := common.GetEthClient(cfg.ETHConfig)
	if err != nil {
		return nil, err
	}
	bridgeSdk, err := http.New(cfg.BridgeConfig.BridgeRpcAddr, "/websocket")
	if err != nil {
		log.Errorf("startServer - failed to new Tendermint Cli: %v", err)
		return nil, err
	}

	CMPrivk, CMAcc, err := common.GetBridgePrivateKey(cfg.BridgeConfig.BridgeWallet, []byte(cfg.BridgeConfig.BridgeWalletPwd))
	if err != nil {
		return nil, err
	}

	var gasPrice ctypes.DecCoins
	acc := authKeeper.GetAccount(*ctx, CMAcc)
	CMSeq := &CosmosSeq{
		lock: sync.Mutex{},
		val:  acc.GetSequence(),
	}
	if gasPrice, err = ctypes.ParseDecCoins(cfg.BridgeConfig.BridgeGasPrice); err != nil {
		return nil, err
	}
	CMFees, err := common.CalcCosmosFees(gasPrice, cfg.BridgeConfig.BridgeGas)
	if err != nil {
		return nil, err
	}
	CMGas := cfg.BridgeConfig.BridgeGas

	contractabi, err := abi.JSON(strings.NewReader(pbridge_abi.PBridgeABI))
	if err != nil {
		return nil, err
	}
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	ks := common.NewEthKeyStore(cfg.ETHConfig, chainId)
	accArr := ks.GetAccounts()
	if len(cfg.ETHConfig.KeyStorePwdSet) == 0 {
		fmt.Println("please input the passwords for ethereum keystore: ")
	}
	if err = ks.UnlockKeys(cfg.ETHConfig); err != nil {
		return nil, err
	}
	senders := make([]*EthSender, len(accArr))
	for i, v := range senders {
		v = &EthSender{}
		v.acc = accArr[i]
		v.ethClient = ethClient
		v.keyStore = ks
		v.config = cfg
		v.contractAbi = &contractabi
		v.nonceManager = common.NewNonceManager(ethClient)
		v.cmap = make(map[string]chan *EthTxInfo)

		senders[i] = v
	}

	em := &EthManager{
		cfg:          cfg,
		restClient:   restClient,
		ethClient:    ethClient,
		senders:      senders,
		CMAcc:        CMAcc,
		CMPrivk:      CMPrivk,
		CMSeq:        CMSeq,
		CMAccNum:     acc.GetAccountNumber(),
		CMFees:       CMFees,
		CMGas:        CMGas,
		bridgeClient: bridgeSdk,
		cdc:          cdc,
	}
	return em, nil
}

// SyncHeader fetch block headers and locked deposit events from ethereum chain
// curHeight: current eth height in pi-bridge
func (em *EthManager) SyncHeadersAndEvents(curHeight uint64) {
	cfg := em.cfg.ETHConfig
	ccMsgs := make([]*trxtypes.MsgCreateTrx, 0)
	var headerSync []string
	// 1. get eth lastest height
	latestHeight, err := common.GetEthHeight(cfg.RestURL, em.restClient)
	if err != nil {
		log.Errorf("SyncHeadersAndEvents - cannot get eth node height, err: %s", err)
		return
	}
	if latestHeight-curHeight <= config.ETH_USEFUL_BLOCK_NUM {
		return
	}
	// 2. fetch eth blocks
	height := curHeight + 1
	for ; height <= latestHeight-config.ETH_USEFUL_BLOCK_NUM; height++ {
		hdr := em.GetHeader(height)
		if hdr == "" {
			break
		}
		headerSync = append(headerSync, hdr)
		if err != nil {
			break
		}
		// get locked deposit events
		ccMsgs, err = em.fetchCrossOutFunds(height)
		if err != nil {
			log.Errorf("SyncHeadersAndEvents - cannot fetch lock deposit events, err: %s", err)
			return
		}
		// batch control
		if len(headerSync) >= cfg.HeadersPerBatch {
			break
		}
	}
	for _, v := range headerSync {
		fmt.Println(v)
	}
	for _, v := range ccMsgs {
		fmt.Println(v.String())
	}
	// 3. commit headers to pi bridge
	if len(headerSync) > 0 {
		index := strconv.FormatUint(cfg.SideChainId, 10)
		msg := blockmoduletypes.NewMsgCreateBlock(em.CMAcc.String(), index, int32(cfg.SideChainId), "", headerSync)
		_, err := em.commitCosmosTx([]ctypes.Msg{msg})
		if err != nil {
			log.Errorf("SyncHeadersAndEvents - commit headers message: %s, err: %s", msg.String(), err)
			return
		}
	}
	// 4. commit locked deposit events to pi bridge
	if len(ccMsgs) > 0 {
		for _, msg := range ccMsgs {
			_, err := em.commitCosmosTx([]ctypes.Msg{msg})
			if err != nil {
				log.Errorf("SyncHeadersAndEvents - commit locked deposit events message: %s, err: %s", msg.String(), err)
				return
			}
		}
	}
}

func (em *EthManager) GetLatestHight() uint64 {
	cfg := em.cfg.ETHConfig
	latestHeight, err := common.GetEthHeight(cfg.RestURL, em.restClient)
	if err != nil {
		log.Errorf("EthManager GetLastestHight - cannot get node height, err: %s", err)
		return 0
	}
	return latestHeight
}

func (em *EthManager) GetHeader(height uint64) string {
	hdr, err := em.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(height)))
	if err != nil {
		log.Errorf("handleBlockHeader - GetNodeHeader on height :%d failed", height)
		return ""
	}
	rawHdr, _ := hdr.MarshalJSON()
	return hex.EncodeToString(rawHdr)
}

func (em *EthManager) commitCosmosTx(msgs []ctypes.Msg) (string, error) {
	seq := em.CMSeq.GetAndAdd()
	encCfg := simapp.MakeTestEncodingConfig()
	// txConfig := authtx.NewTxConfig(em.cdc, authtx.DefaultSignModes)
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		return "", err
	}
	txBuilder.SetGasLimit(em.CMGas)
	txBuilder.SetFeeAmount(em.CMFees)

	signerData := xauthsigning.SignerData{
		// ChainID:       chainID,
		AccountNumber: em.CMAccNum,
		Sequence:      seq,
	}
	sigV2, err := tx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, em.CMPrivk, encCfg.TxConfig, seq)
	if err != nil {
		return "", err
	}
	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return "", err
	}

	// Generate a JSON string.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return "", err
	}
	err = em.sendTx(context.Background(), txBytes)
	return "", err
}

func (em *EthManager) sendTx(ctx context.Context, txBytes []byte) error {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		em.cfg.BridgeConfig.BridgeRpcAddr,
		// "127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	txClient := typestx.NewServiceClient(grpcConn)
	grpcRes, err := txClient.BroadcastTx(
		ctx,
		&typestx.BroadcastTxRequest{
			Mode:    typestx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(grpcRes.TxResponse.Code) // Should be `0` if the tx is successful

	return nil
}

// fetch cross chain tx
func (em *EthManager) fetchCrossOutFunds(height uint64) ([]*trxtypes.MsgCreateTrx, error) {
	cfg := em.cfg.ETHConfig
	ccMsgs := make([]*trxtypes.MsgCreateTrx, 0)
	lockAddress := ethcommon.HexToAddress(cfg.BridgeContractAddress)
	bridgeContract, err := pbridge_abi.NewPBridge(lockAddress, em.ethClient)
	if err != nil {
		return nil, err
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	events, err := bridgeContract.FilterCrossOutFunds(opt)
	if err != nil {
		log.Errorf("fetchCrossOutFunds - FilterCrossChainEvent error :%s", err.Error())
		return nil, err
	}
	if events == nil {
		log.Infof("fetchCrossOutFunds - no events found on FilterCrossChainEvent")
		return nil, err
	}

	for events.Next() {
		evt := events.Event
		param := &common.MakeTxParam{}
		_ = param.Deserialization(types.NewZeroCopySource([]byte(evt.Raw.Data)))
		proof, err := em.GetProof(evt.Raw.TxHash.String(), evt.Raw.BlockNumber)
		if err != nil {
			log.Errorf("fetchCrossOutFunds - GetProof error :%s", err.Error())
			return nil, err
		}
		fmt.Println("proof:", hex.EncodeToString(proof))
		msg := trxtypes.NewMsgCreateTrx(em.CMAcc.String(), evt.Raw.TxHash.String(), int32(cfg.SideChainId), "",
			hex.EncodeToString(evt.Raw.Data), hex.EncodeToString(proof))
		ccMsgs = append(ccMsgs, msg)
	}
	return ccMsgs, nil
}

// SyncLockDepositEvents sync lock deposit events by height
func (em *EthManager) SyncLockDepositEvents(height uint64) {
	cfg := em.cfg.ETHConfig
	latestHeight := em.GetLatestHight()
	if latestHeight <= height+cfg.BlockConfig {
		return
	}
	ccMsg, _ := em.fetchCrossOutFunds(height)
	if len(ccMsg) > 0 {
		for _, msg := range ccMsg {
			_, err := em.commitCosmosTx([]ctypes.Msg{msg})
			if err != nil {
				return
			}
		}
	}
}

func (em *EthManager) GetProof(txIndex string, txHeight uint64) ([]byte, error) {
	cfg := em.cfg.ETHConfig
	heightHex := hexutil.EncodeBig(big.NewInt(int64(txHeight)))
	proof, err := common.GetProof(cfg.RestURL, cfg.BridgeContractAddress, txIndex, heightHex, em.restClient)
	if err != nil {
		log.Errorf("GetProof - error :%s\n", err.Error())
	}
	return proof, nil
}

type CrossTransfer struct {
	txIndex string
	txId    []byte
	value   []byte
	toChain uint32
	height  uint64
}

func (ct *CrossTransfer) Serialization(sink *types.ZeroCopySink) {
	sink.WriteString(ct.txIndex)
	sink.WriteVarBytes(ct.txId)
	sink.WriteVarBytes(ct.value)
	sink.WriteUint32(ct.toChain)
	sink.WriteUint64(ct.height)
}

func (ct *CrossTransfer) Deserialization(source *types.ZeroCopySource) error {
	txIndex, eof := source.NextString()
	if eof {
		return fmt.Errorf("waiting deserialize txIndex error")
	}
	txId, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("waiting deserialize txId error")
	}
	value, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("waiting deserialize value error")
	}
	toChain, eof := source.NextUint32()
	if eof {
		return fmt.Errorf("waiting deserialize toChain error")
	}
	height, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("waiting deserialize height error")
	}
	ct.txIndex = txIndex
	ct.txId = txId
	ct.value = value
	ct.toChain = toChain
	ct.height = height
	return nil
}

// relay pi-bridge cross chain tx to eth
func (em *EthManager) RelayTx(header *BridgeHeader, txs []*BridgeTx) bool {
	for _, tx := range txs {
		sender := em.selectSender()
		// log.Infof("sender %s is handling bridge tx ( hash: %s, height: %d )",
		// 	sender.acc.Address.String(), tx.Tx.Hash, height)
		if sender.commitDepositEventsWithHeader(header, tx) {
			return false
		}
	}
	return true
}

func (bm *EthManager) selectSender() *EthSender {
	sum := big.NewInt(0)
	balArr := make([]*big.Int, len(bm.senders))
	for i, v := range bm.senders {
	RETRY:
		bal, err := v.Balance()
		if err != nil {
			log.Errorf("failed to get balance for %s: %v", v.acc.Address.String(), err)
			time.Sleep(time.Second)
			goto RETRY
		}
		sum.Add(sum, bal)
		balArr[i] = big.NewInt(sum.Int64())
	}
	sum.Rand(rand.New(rand.NewSource(time.Now().Unix())), sum)
	for i, v := range balArr {
		res := v.Cmp(sum)
		if res == 1 || res == 0 {
			return bm.senders[i]
		}
	}
	return bm.senders[0]
}

type BridgeHeader struct {
	Header  tdmtypes.Header
	Commit  *tdmtypes.Commit
	Valsets []*tdmtypes.Validator
}

type BridgeTx struct {
	Tx          *tdmrpctypes.ResultTx
	ProofHeight int64
	Proof       []byte
	PVal        []byte
	ChainID     uint64
}
