package eth

import (
	"context"
	"fmt"
	"math/big"
	"runtime/debug"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pchain-org/pi-bridge/tools/config"
	"github.com/pkg/errors"
)

var (
	addrs        []*AddrInfo
	privkHexList = []string{
		"b3cc8192cda1532fa013dea49974a6dd722cc1321bab11e0aa860f04b436387d", //0x7e9f34471c71858359E5574f0eDeb03dCa9F5f43
		"0d507211e7dfbab3504134a8fc6e31b5fb44d420e302b936f798f52b15021d55", //0xDD9Cd206BC670d8d131E4b95a6Bd916Ad4338570
		"80bb7fbf5084610f4f9dabf5b3cfc27eee00e86424f20e65cdefb187a3225a58", //0x8d505f08E421d59794F462FF0Cc5b01787838CE0
		"f61955911f7cb7304a182885c2b18d8fc433f23b80ad68a77b1a3d38f94b2c78", //0xCdD69c899028A0de95F5518bA5D2a8FfF7cd7799
	}
	erc20ContractAddressHex = "0x2F1dC4AEb25d882e0823f9DCa31172A9f8Ee9411"
	// txKey                   = "8732bff38f474d3ab41c29c53ae0ff73955460582f47efd74a46b9734c4301c9"
)

func TestCommitMultiSigWithdraw(t *testing.T) {
	// 1.prepare
	var signs []byte
	ctx := context.Background()
	servConfig = config.NewServiceConfig("./../../../cs_config.json")
	FailOnFlag(t, servConfig == nil, "load config error")
	em, err := NewEthManager(nil, nil, servConfig, nil)
	if err != nil {
		fmt.Println(err)
	}
	es := em.selectSender()

	contractAddress := common.HexToAddress(servConfig.ETHConfig.BridgeContractAddress)
	bal, err := es.Balance()
	FailOnErr(t, err, "get sender balance error")
	fmt.Println("sender balance:", bal.String())

	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()
	fmt.Println("random tx:", randomTxKey)
	for _, v := range privkHexList {
		addr := GenAddr(v)
		addrs = append(addrs, addr)
	}
	a1, a2, a3 := addrs[1], addrs[2], addrs[3]
	es.acc = a2.ToAccount()
	FailOnErr(t, err, "import a2 into keystore error")
	for _, v := range addrs {
		bal, err := em.ethClient.BalanceAt(ctx, v.ToAddress(), nil)
		FailOnErr(t, err, "get balance error")
		fmt.Printf("%s Balance: %v\n", v.Address, bal)
	}
	preBal, err := em.ethClient.BalanceAt(ctx, contractAddress, nil)
	FailOnErr(t, err, "get balance error")
	fmt.Printf("contract %s Balance: %v\n", contractAddress.Hex(), preBal)

	// transfer to a1, singed by a2 and a3
	destination := a1.Address
	value := big.NewInt(1).Mul(big.NewInt(E18), big.NewInt(1))
	version := big.NewInt(2)
	for _, addr := range []*AddrInfo{a2, a3} {
		s, err := MultiSignExecute(randomTxKey[2:], addr.PrivkHex, false, erc20ContractAddressHex, destination, value, version)
		FailOnErr(t, err, "create sig failed")
		signs = append(signs, s...)
	}
	params := &MultiSigParams{
		TxKey:       randomTxKey[2:],
		Destination: destination,
		Value:       value,
		IsERC20:     false,
		ERC20Addr:   erc20ContractAddressHex,
		Signs:       signs,
	}
	res := es.commitMultiSigWithdraw(params)
	fmt.Println("send to eth success: ", res)

	// check balance
	time.Sleep(time.Second * 2) // waiting for confirm
	bal, err = em.ethClient.BalanceAt(ctx, contractAddress, nil)
	FailOnErr(t, err, "get balance error")
	fmt.Println("contract balance: ", bal.String())
	tmp := big.NewInt(0)
	FailOnFlag(t, preBal.Cmp(tmp.Add(bal, value)) != 0,
		fmt.Sprintf("contract balance fail, should be %v, got %s", (preBal.Sub(preBal, value)).String(), bal.String()))
	for _, v := range addrs {
		bal, err := em.ethClient.BalanceAt(context.Background(), v.ToAddress(), nil)
		FailOnErr(t, err, "get balance error")
		fmt.Printf("%s Balance: %v\n", v.Address, bal)
	}
}

// MultiSignExecute return sig byte
func MultiSignExecute(txKey, signerPrivkHex string, isERC20 bool, ERC20Addr, destinationAddr string, value, version *big.Int) ([]byte, error) {
	addressTy, _ := abi.NewType("address", "", nil)
	stringTy, _ := abi.NewType("string", "", nil)
	uintTy, _ := abi.NewType("uint256", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)
	arguments := abi.Arguments{
		{
			Type: stringTy,
		}, {
			Type: addressTy,
		}, {
			Type: uintTy,
		}, {
			Type: boolTy,
		}, {
			Type: addressTy,
		}, {
			Type: uintTy,
		},
	}
	bytes, err := arguments.Pack(
		txKey,
		common.HexToAddress(destinationAddr),
		value,
		isERC20,
		common.HexToAddress(ERC20Addr),
		version,
	)
	if err != nil {
		fmt.Println(err)
	}
	hashBytes := crypto.Keccak256(
		bytes,
	)

	privk, err := crypto.HexToECDSA(signerPrivkHex)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(hashBytes, privk)
	if err != nil {
		return nil, errors.Wrap(err, "crypto sign failed")
	}
	return sig, nil
}

// FailOnErr used in testing assert
func FailOnErr(t *testing.T, e error, msg string) {
	if e != nil {
		fmt.Printf("[Fail] on error, %s, %v\n", msg, e)
		debug.PrintStack()
		t.FailNow()
	}
}

// FailOnFlag used in flag testing asset
func FailOnFlag(t *testing.T, flag bool, params ...interface{}) {
	if flag {
		fmt.Printf("[Fail] on flag, %v\n", params)
		debug.PrintStack()
		t.FailNow()
	}
}
