package eth

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	cctypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	ctypes "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/config"
	"github.com/pchain-org/pi-bridge/tools/types"
	blockmoduletypes "github.com/pchain-org/pi-bridge/x/block/types"
	trxtypes "github.com/pchain-org/pi-bridge/x/trx/types"
	"github.com/polynetwork/eth-contracts/go_abi/eccm_abi"
	"github.com/prometheus/common/log"
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
		return fmt.Errorf("Waiting deserialize txIndex error")
	}
	txId, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("Waiting deserialize txId error")
	}
	value, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("Waiting deserialize value error")
	}
	toChain, eof := source.NextUint32()
	if eof {
		return fmt.Errorf("Waiting deserialize toChain error")
	}
	height, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("Waiting deserialize height error")
	}
	ct.txIndex = txIndex
	ct.txId = txId
	ct.value = value
	ct.toChain = toChain
	ct.height = height
	return nil
}
