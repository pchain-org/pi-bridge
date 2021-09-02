package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/config"
	"github.com/pchain-org/pi-bridge/tools/log"
	pbridge_abi "github.com/terry108/cross-chain-contract/pbridge_abi"
)

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

	isSuccess := es.waitTransactionConfirm(hash)
	if isSuccess {
		log.Infof("successful to relay tx to ethereum: (eth_hash: %s, nonce: %d, bridge_hash: %s, eth_explorer: %s)",
			hash.String(), nonce, info.bridgeTxHash, common.GetExplorerUrl(es.keyStore.GetChainId())+hash.String())
	} else {
		log.Errorf("failed to relay tx to ethereum: (eth_hash: %s, nonce: %d, bridge_hash: %s, eth_explorer: %s)",
			hash.String(), nonce, info.bridgeTxHash, common.GetExplorerUrl(es.keyStore.GetChainId())+hash.String())
	}
	return nil
}

// TODO: check the status of tx
func (es *EthSender) waitTransactionConfirm(hash ethcommon.Hash) bool {
	for {
		time.Sleep(time.Second * 1)
		_, ispending, err := es.ethClient.TransactionByHash(context.Background(), hash)
		if err != nil {
			continue
		}
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

// send cross chain tx fo eth
func (es *EthSender) commitDepositEventsWithHeader(header *BridgeHeader, tx *BridgeTx) bool {
	// 1. check the tx status
	eccdAddr := ethcommon.HexToAddress(es.config.ETHConfig.ECCDContractAddress)
	pbridgeAbi, err := pbridge_abi.NewPBridge(eccdAddr, es.ethClient)
	if err != nil {
		panic(fmt.Errorf("failed to new eccm: %v", err))
	}
	completed, err := pbridgeAbi.IsCompletedTx(nil, string(tx.Tx.Hash))
	if err != nil {
		return false
	}
	if completed {
		return true
	}

	// 2. build tx
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
	contractaddr := ethcommon.HexToAddress(es.config.ETHConfig.BridgeContractAddress)
	callMsg := ethereum.CallMsg{
		From: es.acc.Address, To: &contractaddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := es.ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Errorf("commitDepositEventsWithHeader - estimate gas limit error: %s", err.Error())
		return false
	}

	// 3. send tx
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

// TxParams 交易参数
type MultiSigParams struct {
	TxKey       string   // cross chain transaction key
	Destination string   // to address
	Value       *big.Int // cross chain amount
	IsERC20     bool     // is ERC20
	ERC20Addr   string   // ERC20 address
	Signs       []byte   // 多重签名
}

// send cross chain tx fo eth
func (es *EthSender) commitMultiSigWithdraw(msp *MultiSigParams) bool {
	// 1. build tx
	txData, err := es.contractAbi.Pack("createOrSignWithdraw", msp.TxKey, ethcommon.HexToAddress(msp.Destination),
		msp.Value, msp.IsERC20, ethcommon.HexToAddress(msp.ERC20Addr), msp.Signs)
	if err != nil {
		log.Errorf("commitMultiSigWithdraw - err:" + err.Error())
		return false
	}

	gasPrice, err := es.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("commitMultiSigWithdraw - get suggest sas price failed error: %s", err.Error())
		return false
	}
	contractaddr := ethcommon.HexToAddress(es.config.ETHConfig.BridgeContractAddress)
	callMsg := ethereum.CallMsg{
		From: es.acc.Address, To: &contractaddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := es.ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Errorf("commitMultiSigWithdraw - estimate gas limit error: %s", err.Error())
		return false
	}
	fmt.Println("gasLimit:", gasLimit)

	// 3. send tx
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
		bridgeTxHash: msp.TxKey,
	}
	return true
}

func (es *EthSender) Balance() (*big.Int, error) {
	balance, err := es.ethClient.BalanceAt(context.Background(), es.acc.Address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
