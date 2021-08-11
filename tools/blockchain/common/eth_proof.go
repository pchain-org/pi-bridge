package common

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	ethcomm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/tendermint/tendermint/libs/log"
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

type Proof struct {
	AssetAddress string
	FromAddress  string
	ToChainID    uint64
	ToAddress    string
	Args         []byte
}

type StorageProof struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Proof []string `json:"proof"`
}

type ProofAccount struct {
	Nounce   *big.Int
	Balance  *big.Int
	Storage  ethcomm.Hash
	Codehash ethcomm.Hash
}

type ETHProof struct {
	Address       string         `json:"address"`
	Balance       string         `json:"balance"`
	CodeHash      string         `json:"codeHash"`
	Nonce         string         `json:"nonce"`
	StorageHash   string         `json:"storageHash"`
	AccountProof  []string       `json:"accountProof"`
	StorageProofs []StorageProof `json:"storageProof"`
}

type proofReq struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      uint          `json:"id"`
}

type proofRsp struct {
	JsonRPC string     `json:"jsonrpc"`
	Result  ETHProof   `json:"result,omitempty"`
	Error   *jsonError `json:"error,omitempty"`
	Id      uint       `json:"id"`
}

func (ep *ETHProof) String() string {
	bs := bytes.NewBuffer([]byte("ETHProof:\n"))
	bs.WriteString("AccountProof:\n")
	for _, a := range ep.AccountProof {
		bs.WriteString(a + "\n")
	}
	bs.WriteString("Address:")
	bs.WriteString(ep.Address + "\n")
	bs.WriteString("StorageProof:\n")
	for _, s := range ep.StorageProofs {
		bs.WriteString(s.Key + "\n")
		bs.WriteString("proofs:\n[")
		bs.WriteString(strings.Join(s.Proof, "\n"))
		bs.WriteString("]\n")

		bs.WriteString(s.Value + "\n")
	}
	return bs.String()
}

func MappingKeyAt(position1 string, position2 string) ([]byte, error) {

	p1, err := hex.DecodeString(position1)
	if err != nil {
		return nil, err
	}

	p2, err := hex.DecodeString(position2)

	if err != nil {
		return nil, err
	}

	key := crypto.Keccak256(ethcomm.LeftPadBytes(p1, 32), ethcomm.LeftPadBytes(p2, 32))

	return key, nil
}

func (p *Proof) Deserialize(raw string) error {
	vals := strings.Split(raw, "#")
	if len(vals) != 6 {
		return fmt.Errorf("error count of proof deserialize")
	}
	p.AssetAddress = vals[0]
	p.FromAddress = vals[1]
	cid, err := strconv.Atoi(vals[2])
	if err != nil {
		return fmt.Errorf("chain id is not correct")
	}
	p.ToChainID = uint64(cid)
	p.ToAddress = vals[3]
	amt := new(big.Int)
	_, b := amt.SetString(vals[4], 10)
	if !b {
		return fmt.Errorf("amount is not correct")
	}
	p.Args = []byte(vals[5])
	//p.Amount = amt
	//decimal, err := strconv.Atoi(vals[5])
	//if err != nil {
	//	return fmt.Errorf("decimal is not correct")
	//}
	//p.Decimal = decimal

	return nil
}

func GetProof(url string, contractAddress string, key string, blockheight string, restClient *RestClient) ([]byte, error) {
	req := &proofReq{
		JsonRPC: "2.0",
		Method:  "eth_getProof",
		Params:  []interface{}{contractAddress, []string{key}, blockheight},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("get_ethproof: marshal req err: %s", err)
	}
	rspdata, err := restClient.SendRestRequest(url, reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetProof: send request err: %s", err)
	}
	rsp := &proofRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetProof, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("GetProof, unmarshal resp err: %s", rsp.Error.Message)
	}
	result, err := json.Marshal(rsp.Result)
	if err != nil {
		return nil, fmt.Errorf("GetProof, Marshal result err: %s", err)
	}
	//fmt.Printf("proof res is:%s\n", string(result))
	return result, nil
}

// used by quorum
func VerifyMerkleProofLegacy(ethProof *ETHProof, blockData *types.Header, contractAddr []byte) ([]byte, error) {
	return VerifyMerkleProof(ethProof, To1559(blockData), contractAddr)
}

func VerifyMerkleProof(ethProof *ETHProof, blockData *Header, contractAddr []byte) ([]byte, error) {
	//1. prepare verify account
	nodeList := new(light.NodeList)

	for _, s := range ethProof.AccountProof {
		p := Replace0x(s)
		nodeList.Put(nil, ethcomm.Hex2Bytes(p))
	}
	ns := nodeList.NodeSet()

	addr := ethcomm.Hex2Bytes(Replace0x(ethProof.Address))
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

	storageHash := ethcomm.HexToHash(Replace0x(ethProof.StorageHash))
	codeHash := ethcomm.HexToHash(Replace0x(ethProof.CodeHash))

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
	storageKey := crypto.Keccak256(ethcomm.HexToHash(Replace0x(sp.Key)).Bytes())

	for _, prf := range sp.Proof {
		nodeList.Put(nil, ethcomm.Hex2Bytes(Replace0x(prf)))
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
		log.NewNopLogger().Error("checkProofResult, rlp.DecodeBytes error:%s\n", err)
		return false
	}

	var s []byte
	for i := len(s_temp); i < 32; i++ {
		s = append(s, 0)
	}
	s = append(s, s_temp...)
	hash := crypto.Keccak256(value)
	return bytes.Equal(s, hash)
}
