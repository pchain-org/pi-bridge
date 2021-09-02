package common

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestGetProof(t *testing.T) {
	url := "http://127.0.0.1:8545"
	restClient := NewRestClient()
	contactAdr := "0x9CE47a47Dc42aCE1539E6b9276C0177Ea7fDd9fB"
	key := "0x8732bff38f474d3ab41c29c53ae0ff73955460582f47efd74a46b9734c4301c9"
	heightHex := hexutil.EncodeBig(big.NewInt(35))
	_, err := GetProof(url, contactAdr, key, heightHex, restClient)
	if err != nil {
		fmt.Println(err)
	}
}
