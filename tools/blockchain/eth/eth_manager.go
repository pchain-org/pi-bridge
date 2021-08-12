package eth

import (
	"context"
	"encoding/hex"
	"errors"
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
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/config"
	"github.com/pchain-org/pi-bridge/tools/log"
	"github.com/pchain-org/pi-bridge/tools/types"
	blockmoduletypes "github.com/pchain-org/pi-bridge/x/block/types"
	trxtypes "github.com/pchain-org/pi-bridge/x/trx/types"
	"github.com/polynetwork/eth-contracts/go_abi/eccd_abi"
	"github.com/polynetwork/eth-contracts/go_abi/eccm_abi"
	tdmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
	tdmtypes "github.com/tendermint/tendermint/types"
)

const (
	ChanLen = 64
)

type EthManager struct {
	cfg           *config.ServiceConfig
	restClient    *common.RestClient
	ethClient     *ethclient.Client
	currentHeight uint64
	forceHeight   uint64

	CMPrivk  cctypes.PrivKey
	CMAcc    ctypes.AccAddress
	CMSeq    *CosmosSeq
	CMAccNum uint64
	CMFees   ctypes.Coins
	CMGas    uint64
	CMCdc    *codec.AminoCodec

	senders []*EthSender
}

type SyncHeader struct {
	latestHeight  uint64
	endHeight     uint64
	headerSync    []string
	signdHeaderTx string
}

type SyncCrossTx struct {
	latestHeight uint64
	endHeight    uint64
	txs          []TxData
}

type TxData struct {
	txSync  string
	signdTx string
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

func NewEthManager(ctx ctypes.Context, cfg *config.ServiceConfig, authKeeper authkeeper.AccountKeeper) (*EthManager, error) {
	restClient := common.NewRestClient()
	ethClient, err := common.GetEthClient(cfg.ETHConfig)
	if err != nil {
		return nil, err
	}
	CMPrivk, CMAcc, err := common.GetBridgePrivateKey(cfg.BridgeConfig.BridgeWallet, []byte(cfg.BridgeConfig.BridgeWalletPwd))
	if err != nil {
		return nil, err
	}

	var gasPrice ctypes.DecCoins
	acc := authKeeper.GetAccount(ctx, CMAcc)
	CMSeq := &CosmosSeq{
		lock: sync.Mutex{},
		val:  acc.GetSequence(),
	}
	CMFees, err := common.CalcCosmosFees(gasPrice, cfg.BridgeConfig.BridgeGas)
	if err != nil {
		return nil, err
	}
	CMGas := cfg.BridgeConfig.BridgeGas

	contractabi, err := abi.JSON(strings.NewReader(eccm_abi.EthCrossChainManagerABI))
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
		cfg:        cfg,
		restClient: restClient,
		ethClient:  ethClient,
		CMAcc:      CMAcc,
		CMPrivk:    CMPrivk,
		CMSeq:      CMSeq,
		CMAccNum:   acc.GetAccountNumber(),
		CMFees:     CMFees,
		CMGas:      CMGas,
	}
	return em, nil
}

// SyncHeader fetch block from ethereum chain
// ethConfig: get from app
// curHeight: current eth height in pi-bridge
func (em *EthManager) SyncHeader(curHeight uint64) (*SyncHeader, error) {
	cfg := em.cfg.ETHConfig
	// 1. get lastest height
	latestHeight, err := common.GetEthHeight(cfg.RestURL, em.restClient)
	if err != nil {
		log.Errorf("SyncHeader - cannot get node height, err: %s", err)
		return nil, err
	}
	syncData := &SyncHeader{
		latestHeight: latestHeight,
	}
	if latestHeight-curHeight <= config.ETH_USEFUL_BLOCK_NUM {
		return syncData, nil
	}
	// 2. fetch blocks and locked deposit events
	height := curHeight + 1
	for ; height < latestHeight-config.ETH_USEFUL_BLOCK_NUM; height++ {
		hdr := em.GetHeader(height)
		if hdr == "" {
			break
		}
		syncData.headerSync = append(syncData.headerSync, hdr)
		if err != nil {
			break
		}
		if len(syncData.headerSync) >= cfg.HeadersPerBatch {
			break
		}
	}
	if len(syncData.headerSync) > 0 {
		index := strconv.FormatUint(cfg.SideChainId, 10)
		msg := blockmoduletypes.NewMsgCreateBlock(em.CMAcc.String(), index, int32(cfg.SideChainId), "", syncData.headerSync)
		signdMsg, err := em.createCosmosTx([]ctypes.Msg{msg})
		if err != nil {
			return syncData, err
		}
		syncData.endHeight += uint64(len(syncData.headerSync))
		syncData.signdHeaderTx = signdMsg
	}
	return syncData, nil
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

func (em *EthManager) createCosmosTx(msgs []ctypes.Msg) (string, error) {
	seq := em.CMSeq.GetAndAdd()
	encCfg := simapp.MakeTestEncodingConfig()
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
	txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return "", err
	}
	txJSON := string(txJSONBytes)
	return txJSON, nil
}

func (em *EthManager) GetLockDepositEvents(height uint64) (*SyncCrossTx, error) {
	cfg := em.cfg.ETHConfig
	latestHeight := em.GetLatestHight()
	syncCrossTx := &SyncCrossTx{
		latestHeight: latestHeight,
		endHeight:    height,
	}
	if latestHeight <= height+cfg.BlockConfig {
		return syncCrossTx, errors.New("OUT_OF_SAFE_BLOCK")
	}
	lockAddress := ethcommon.HexToAddress(cfg.ECCMContractAddress)
	lockContract, err := eccm_abi.NewEthCrossChainManager(lockAddress, em.ethClient)
	if err != nil {
		return syncCrossTx, err
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	events, err := lockContract.FilterCrossChainEvent(opt, nil)
	if err != nil {
		log.Errorf("GetLockDepositEvents - FilterCrossChainEvent error :%s", err.Error())
		return syncCrossTx, err
	}
	if events == nil {
		log.Infof("GetLockDepositEvents - no events found on FilterCrossChainEvent")
		return syncCrossTx, err
	}
	for events.Next() {
		evt := events.Event
		var isTarget bool
		if len(cfg.TargetContracts) > 0 {
			toContractStr := evt.ProxyOrAssetContract.String()
			for _, v := range cfg.TargetContracts {
				toChainIdArr, ok := v[toContractStr]
				if ok {
					if len(toChainIdArr["outbound"]) == 0 {
						isTarget = true
						break
					}
					for _, id := range toChainIdArr["outbound"] {
						if id == evt.ToChainId {
							isTarget = true
							break
						}
					}
					if isTarget {
						break
					}
				}
			}
			if !isTarget {
				continue
			}
		}
		param := &common.MakeTxParam{}
		_ = param.Deserialization(types.NewZeroCopySource([]byte(evt.Rawdata)))
		index := big.NewInt(0)
		index.SetBytes(evt.TxId)
		crossTx := &CrossTransfer{
			txIndex: common.EncodeBigInt(index),
			txId:    evt.Raw.TxHash.Bytes(),
			toChain: uint32(evt.ToChainId),
			value:   []byte(evt.Rawdata),
			height:  height,
		}
		log.Infof("GetLockDepositEvents -  height: %d", height)

		// get proof
		proof, err := em.GetProof(crossTx, height)
		if err != nil {
			log.Errorf("GetLockDepositEvents - GetProof error :%s", err.Error())
			return nil, err
		}
		msg := trxtypes.NewMsgCreateTrx(em.CMAcc.String(), common.EncodeBigInt(index), int32(cfg.SideChainId), "",
			hex.EncodeToString(crossTx.value), hex.EncodeToString(proof))
		signdMsg, err := em.createCosmosTx([]ctypes.Msg{msg})
		if err != nil {
			return nil, err
		}
		txData := TxData{
			txSync:  hex.EncodeToString(crossTx.value),
			signdTx: signdMsg,
		}
		syncCrossTx.txs = append(syncCrossTx.txs, txData)
	}
	return syncCrossTx, nil
}

func (em *EthManager) GetProof(crossTx *CrossTransfer, refHeight uint64) ([]byte, error) {
	cfg := em.cfg.ETHConfig
	key := crossTx.txIndex
	keyBytes, err := common.MappingKeyAt(key, "01")
	if err != nil {
		log.Errorf("handleLockDepositEvents - MappingKeyAt error:%s\n", err.Error())
		return nil, err
	}
	if refHeight <= crossTx.height+cfg.BlockConfig {
		return nil, nil
	}
	height := int64(refHeight - cfg.BlockConfig)
	heightHex := hexutil.EncodeBig(big.NewInt(height))
	proofKey := hexutil.Encode(keyBytes)
	proof, err := common.GetProof(cfg.RestURL, cfg.ECCDContractAddress, proofKey, heightHex, em.restClient)
	if err != nil {
		log.Errorf("handleLockDepositEvents - error :%s\n", err.Error())
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

type EthSender struct {
	acc          accounts.Account
	keyStore     *common.EthKeyStore
	cmap         map[string]chan *EthTxInfo
	nonceManager *common.NonceManager
	ethClient    *ethclient.Client
	config       *config.ServiceConfig
	contractAbi  *abi.ABI
}

type EthTxInfo struct {
	txData       []byte
	gasLimit     uint64
	gasPrice     *big.Int
	contractAddr ethcommon.Address
	bridgeTxHash string
}

func (es *EthSender) Balance() (*big.Int, error) {
	balance, err := es.ethClient.BalanceAt(context.Background(), es.acc.Address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (es *EthSender) getRouter() string {
	return strconv.FormatInt(rand.Int63n(es.config.RoutineNum), 10)
}

func (es *EthSender) sendTxToEth(info *EthTxInfo) error {
	nonce := es.nonceManager.GetAddressNonce(es.acc.Address)
	tx := ethtypes.NewTransaction(nonce, info.contractAddr, big.NewInt(0), info.gasLimit, info.gasPrice, info.txData)
	signedtx, err := es.keyStore.SignTransaction(tx, es.acc)
	if err != nil {
		es.nonceManager.ReturnNonce(es.acc.Address, nonce)
		return fmt.Errorf("commitDepositEventsWithHeader - sign raw tx error and return nonce %d: %v", nonce, err)
	}
	err = es.ethClient.SendTransaction(context.Background(), signedtx)
	if err != nil {
		es.nonceManager.ReturnNonce(es.acc.Address, nonce)
		return fmt.Errorf("commitDepositEventsWithHeader - send transaction error and return nonce %d: %v", nonce, err)
	}
	hash := signedtx.Hash()

	isSuccess := es.waitTransactionConfirm(info.bridgeTxHash, hash)
	if isSuccess {
		log.Infof("successful to relay tx to ethereum: (eth_hash: %s, nonce: %d, poly_hash: %s, eth_explorer: %s)",
			hash.String(), nonce, info.bridgeTxHash, common.GetExplorerUrl(es.keyStore.GetChainId())+hash.String())
	} else {
		log.Errorf("failed to relay tx to ethereum: (eth_hash: %s, nonce: %d, poly_hash: %s, eth_explorer: %s)",
			hash.String(), nonce, info.bridgeTxHash, common.GetExplorerUrl(es.keyStore.GetChainId())+hash.String())
	}
	return nil
}

// TODO: check the status of tx
func (es *EthSender) waitTransactionConfirm(polyTxHash string, hash ethcommon.Hash) bool {
	for {
		time.Sleep(time.Second * 1)
		_, ispending, err := es.ethClient.TransactionByHash(context.Background(), hash)
		if err != nil {
			continue
		}
		log.Debugf("( eth_transaction %s, poly_tx %s ) is pending: %v", hash.String(), polyTxHash, ispending)
		if ispending {
			continue
		} else {
			receipt, err := es.ethClient.TransactionReceipt(context.Background(), hash)
			if err != nil {
				continue
			}
			return receipt.Status == ethtypes.ReceiptStatusSuccessful
		}
	}
}

func (es *EthSender) commitDepositEventsWithHeader(header *BridgeHeader, tx *BridgeTx) bool {
	eccdAddr := ethcommon.HexToAddress(es.config.ETHConfig.ECCDContractAddress)
	eccd, err := eccd_abi.NewEthCrossChainData(eccdAddr, es.ethClient)
	if err != nil {
		panic(fmt.Errorf("failed to new eccm: %v", err))
	}
	// TODO check tx
	fromTx := [32]byte{}
	// copy(fromTx[:], param.TxHash[:32])
	res, _ := eccd.CheckIfFromChainTxExist(nil, tx.ChainID, fromTx)
	if res {
		// log.Debugf("already relayed to eth: ( from_chain_id: %d, from_txhash: %x,  param.Txhash: %x)",
		// 	param.FromChainID, param.TxHash, param.MakeTxParam.TxHash)
		return true
	}

	headerdata := header.Header
	txData, err := es.contractAbi.Pack("verifyHeaderAndExecuteTx", headerdata, header.Valsets, header.Valsets)
	if err != nil {
		log.Errorf("commitDepositEventsWithHeader - err:" + err.Error())
		return false
	}

	gasPrice, err := es.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("commitDepositEventsWithHeader - get suggest sas price failed error: %s", err.Error())
		return false
	}
	contractaddr := ethcommon.HexToAddress(es.config.ETHConfig.ECCMContractAddress)
	callMsg := ethereum.CallMsg{
		From: es.acc.Address, To: &contractaddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := es.ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Errorf("commitDepositEventsWithHeader - estimate gas limit error: %s", err.Error())
		return false
	}

	k := es.getRouter()
	c, ok := es.cmap[k]
	if !ok {
		c = make(chan *EthTxInfo, ChanLen)
		es.cmap[k] = c
		go func() {
			for v := range c {
				if err = es.sendTxToEth(v); err != nil {
					log.Errorf("failed to send tx to ethereum: error: %v, txData: %s", err, hex.EncodeToString(v.txData))
				}
			}
		}()
	}
	//TODO: could be blocked
	c <- &EthTxInfo{
		txData:       txData,
		contractAddr: contractaddr,
		gasPrice:     gasPrice,
		gasLimit:     gasLimit,
		bridgeTxHash: string(tx.Tx.Hash),
	}
	return true
}
