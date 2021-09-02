package eth

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/config"

	pbridge_abi "github.com/terry108/cross-chain-contract/pbridge_abi"
)

var servConfig *config.ServiceConfig

// func init() {

// 	servConfig = config.NewServiceConfig("./../../../cs_config.json")
// 	if servConfig == nil {
// 		fmt.Print("test init error")
// 	}
// 	em, err := NewEthManager(nil, nil, servConfig, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	_ = em
// }

func TestMinManagers(t *testing.T) {
	servConfig = config.NewServiceConfig("./../../../cs_config.json")
	if servConfig == nil {
		fmt.Print("test init error")
	}
	em, err := NewEthManager(nil, nil, servConfig, nil)
	if err != nil {
		fmt.Println(err)
	}
	cfg := em.cfg.ETHConfig
	lockAddress := ethcommon.HexToAddress(cfg.BridgeContractAddress)
	pbridgeAbi, err := pbridge_abi.NewPBridge(lockAddress, em.ethClient)
	if err != nil {
		fmt.Println(err)
	}
	res, err := pbridgeAbi.MinManagers(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.String())
}

func TestFetchCrossOutFunds(t *testing.T) {
	servConfig = config.NewServiceConfig("./../../../cs_config.json")
	if servConfig == nil {
		fmt.Print("test init error")
	}
	em, err := NewEthManager(nil, nil, servConfig, nil)
	if err != nil {
		fmt.Println(err)
	}

	msgs, err := em.fetchCrossOutFunds(35)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range msgs {
		fmt.Println(v.String())
	}
}

func TestNewCrossChain(t *testing.T) {
	// cfg := em.cfg.ETHConfig
	// lockAddress := ethcommon.HexToAddress(cfg.BridgeContractAddress)
	// // lockContract, err := eccm_abi.NewEthCrossChainManager(lockAddress, em.ethClient)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// opt := bind.TransactOpts{}
	// toChainID := uint64(100)
	// toContact := []byte("test")
	// method := []byte("test")
	// txData := []byte("test")
	// fmt.Println(txData)
	// tx, err := lockContract.CrossChain(&opt, toChainID, toContact, method, txData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(tx.Hash().String())
}

func TestSendEthTx(t *testing.T) {
	servConfig = config.NewServiceConfig("./../../../cs_config.json")
	if servConfig == nil {
		fmt.Print("test init error")
	}
	em, err := NewEthManager(nil, nil, servConfig, nil)
	if err != nil {
		fmt.Println(err)
	}

	es := em.selectSender()
	var txData []byte
	// txData, err := es.contractAbi.Pack("verifyHeaderAndExecuteTx", headerdata, header.Valsets, header.Valsets)
	if err != nil {
		fmt.Printf("commitDepositEventsWithHeader - err:" + err.Error())
		return
	}
	gasPrice, err := es.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("commitDepositEventsWithHeader - get suggest sas price failed error: %s \n", err.Error())
		return
	}
	contractaddr := ethcommon.HexToAddress(es.config.ETHConfig.BridgeContractAddress)
	callMsg := ethereum.CallMsg{
		From: es.acc.Address, To: &contractaddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := es.ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		fmt.Printf("commitDepositEventsWithHeader - estimate gas limit error: %s \n", err.Error())
		return
	}

	eti := &EthTxInfo{
		txData:       txData,
		contractAddr: contractaddr,
		gasPrice:     gasPrice,
		gasLimit:     gasLimit,
	}

	err = es.sendTxToEth(eti)
	if err != nil {
		fmt.Printf("commitDepositEventsWithHeader - send tx error: %s \n", err.Error())
	}
}

func TestKeyStore(t *testing.T) {
	servConfig = config.NewServiceConfig("./../../../cs_config.json")
	if servConfig == nil {
		fmt.Print("test init error")
	}
	em, err := NewEthManager(nil, nil, servConfig, nil)
	if err != nil {
		fmt.Println(err)
	}

	ethClient, err := ethclient.Dial(em.cfg.ETHConfig.RestURL)
	if err != nil {
		fmt.Printf("GetEthClient - cannot dial sync node, err: %s", err)
	}
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	ks := common.NewEthKeyStore(em.cfg.ETHConfig, chainId)
	accArr := ks.GetAccounts()
	if len(em.cfg.ETHConfig.KeyStorePwdSet) == 0 {
		fmt.Println("please input the passwords for ethereum keystore: ")
	}
	if err = ks.UnlockKeys(em.cfg.ETHConfig); err != nil {
		fmt.Println(err)
	}

	fmt.Println(accArr)
}
