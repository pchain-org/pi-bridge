package common

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/prometheus/common/log"

	ecom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

const (
	// source from https://github.com/ethereum/go-ethereum/blob/master/consensus/ethash/consensus.go#L45
	AllowedFutureBlockTime = 15 * time.Second // Max time from current time allowed for blocks, before they're considered future blocks
	EpochLength            = 30000
	MaxEpoch               = 2048
	DatasetInitBytes       = 1 << 30
	DatasetGrowthBytes     = 1 << 23
	MixBytes               = 128
	LoopAccesses           = 64
	HashBytes              = 64
	HashWords              = 16
	DatasetParents         = 256
	CacheInitBytes         = 1 << 24
	CacheGrowthBytes       = 1 << 17
	CacheRounds            = 3
)

// used by quorum
func VerifyMerkleProofLegacy(ethProof *ETHProof, blockData *types.Header, contractAddr []byte) ([]byte, error) {
	return VerifyMerkleProof(ethProof, To1559(blockData), contractAddr)
}

func VerifyMerkleProof(ethProof *ETHProof, blockData *Header, contractAddr []byte) ([]byte, error) {
	//1. prepare verify account
	nodeList := new(light.NodeList)

	for _, s := range ethProof.AccountProof {
		p := Replace0x(s)
		nodeList.Put(nil, ecom.Hex2Bytes(p))
	}
	ns := nodeList.NodeSet()

	addr := ecom.Hex2Bytes(Replace0x(ethProof.Address))
	if !bytes.Equal(addr, contractAddr) {
		return nil, fmt.Errorf("verifyMerkleProof, contract address is error, proof address: %s, side chain address: %s", ethProof.Address, hex.EncodeToString(contractAddr))
	}
	acctKey := crypto.Keccak256(addr)

	// 2. verify account proof
	acctVal, err := trie.VerifyProof(blockData.Root, acctKey, ns)
	if err != nil {
		return nil, fmt.Errorf("verifyMerkleProof, verify account proof error:%s\n", err)
	}

	nounce := new(big.Int)
	_, ok := nounce.SetString(Replace0x(ethProof.Nonce), 16)
	if !ok {
		return nil, fmt.Errorf("verifyMerkleProof, invalid format of nounce:%s\n", ethProof.Nonce)
	}

	balance := new(big.Int)
	_, ok = balance.SetString(Replace0x(ethProof.Balance), 16)
	if !ok {
		return nil, fmt.Errorf("verifyMerkleProof, invalid format of balance:%s\n", ethProof.Balance)
	}

	storageHash := ecom.HexToHash(Replace0x(ethProof.StorageHash))
	codeHash := ecom.HexToHash(Replace0x(ethProof.CodeHash))

	acct := &ProofAccount{
		Nounce:   nounce,
		Balance:  balance,
		Storage:  storageHash,
		Codehash: codeHash,
	}

	acctrlp, err := rlp.EncodeToBytes(acct)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(acctrlp, acctVal) {
		return nil, fmt.Errorf("verifyMerkleProof, verify account proof failed, wanted:%v, get:%v", acctrlp, acctVal)
	}

	//3.verify storage proof
	nodeList = new(light.NodeList)
	if len(ethProof.StorageProofs) != 1 {
		return nil, fmt.Errorf("verifyMerkleProof, invalid storage proof format")
	}

	sp := ethProof.StorageProofs[0]
	storageKey := crypto.Keccak256(ecom.HexToHash(Replace0x(sp.Key)).Bytes())

	for _, prf := range sp.Proof {
		nodeList.Put(nil, ecom.Hex2Bytes(Replace0x(prf)))
	}

	ns = nodeList.NodeSet()
	val, err := trie.VerifyProof(storageHash, storageKey, ns)
	if err != nil {
		return nil, fmt.Errorf("verifyMerkleProof, verify storage proof error:%s\n", err)
	}

	return val, nil
}

func CheckProofResult(result, value []byte) bool {
	var s_temp []byte
	err := rlp.DecodeBytes(result, &s_temp)
	if err != nil {
		log.Errorf("checkProofResult, rlp.DecodeBytes error:%s\n", err)
		return false
	}
	//
	var s []byte
	for i := len(s_temp); i < 32; i++ {
		s = append(s, 0)
	}
	s = append(s, s_temp...)
	hash := crypto.Keccak256(value)
	return bytes.Equal(s, hash)
}
