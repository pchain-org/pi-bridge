package common

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/crypto"
	ctypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
)

func GetBridgePrivateKey(path string, pwd []byte) (ctypes.PrivKey, types.AccAddress, error) {
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, types.AccAddress{}, err
	}
	privKey, _, err := crypto.UnarmorDecryptPrivKey(string(bz), string(pwd))
	if err != nil {
		return nil, types.AccAddress{}, fmt.Errorf("failed to decrypt private key: v", err)
	}

	return privKey, types.AccAddress(privKey.PubKey().Address().Bytes()), nil
}

func CalcCosmosFees(gasPrice types.DecCoins, gas uint64) (types.Coins, error) {
	if gasPrice.IsZero() {
		return types.Coins{}, errors.New("gas price is zero")
	}
	if gas == 0 {
		return types.Coins{}, errors.New("gas is zero")
	}
	glDec := types.NewDec(int64(gas))
	fees := make(types.Coins, len(gasPrice))
	for i, gp := range gasPrice {
		fee := gp.Amount.Mul(glDec)
		fees[i] = types.NewCoin(gp.Denom, fee.Ceil().RoundInt())
	}
	return fees, nil
}

// func NewCodecForRelayer() *codec.AminoCodec {
// 	cdc := codec.AminoCodec{
// 		LegacyAmino: codec.NewLegacyAmino(),
// 	}
// 	bank
// 	bank.RegisterCodec(cdc)
// 	types.RegisterCodec(cdc)
// 	codec.RegisterCrypto(cdc)
// 	auth.RegisterCodec(cdc)
// 	// btcx.RegisterCodec(cdc)
// 	// ccm.RegisterCodec(cdc)
// 	// ft.RegisterCodec(cdc)
// 	// headersync.RegisterCodec(cdc)
// 	// lockproxy.RegisterCodec(cdc)
// 	return cdc
// }
