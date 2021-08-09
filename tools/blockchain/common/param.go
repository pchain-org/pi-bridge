package common

import (
	"strings"
)

var (
	KEY_PREFIX_BTC = "btc"

	KEY_PREFIX_BTC_VOTE = "btcVote"
	REQUEST             = "request"
	DONE_TX             = "doneTx"

	NOTIFY_MAKE_PROOF = "makeProof"
)

type InitRedeemScriptParam struct {
	RedeemScript string
}

type MultiSignParam struct {
	ChainID   uint64
	RedeemKey string
	TxHash    []byte
	Address   string
	Signs     [][]byte
}

type ToMerkleValue struct {
	TxHash      []byte
	FromChainID uint64
	MakeTxParam *MakeTxParam
}

func Replace0x(s string) string {
	return strings.Replace(strings.ToLower(s), "0x", "", 1)
}
