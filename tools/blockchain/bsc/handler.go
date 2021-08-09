package bsc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	ecommon "github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Handler ...
type Handler struct {
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

var (
	inMemoryHeaders = 400
	inMemoryGenesis = 40
	extraVanity     = 32                        // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal       = crypto.SignatureLength    // Fixed number of extra-data suffix bytes reserved for signer seal
	uncleHash       = etypes.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
	diffInTurn      = big.NewInt(2)             // Block difficulty for in-turn signatures
	diffNoTurn      = big.NewInt(1)             // Block difficulty for out-of-turn signatures

	GasLimitBoundDivisor uint64 = 256 // The bound divisor of the gas limit, used in update calculations.
)

// Proof ...
type Proof struct {
	Address       string         `json:"address"`
	Balance       string         `json:"balance"`
	CodeHash      string         `json:"codeHash"`
	Nonce         string         `json:"nonce"`
	StorageHash   string         `json:"storageHash"`
	AccountProof  []string       `json:"accountProof"`
	StorageProofs []StorageProof `json:"storageProof"`
}

// StorageProof ...
type StorageProof struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Proof []string `json:"proof"`
}

// ProofAccount ...
type ProofAccount struct {
	Nounce   *big.Int
	Balance  *big.Int
	Storage  ecommon.Hash
	Codehash ecommon.Hash
}

func verifyMerkleProof(bscProof *Proof, blockData *etypes.Header, contractAddr []byte) ([]byte, error) {
	//1. prepare verify account
	nodeList := new(light.NodeList)

	for _, s := range bscProof.AccountProof {
		p := common.Replace0x(s)
		nodeList.Put(nil, ecommon.Hex2Bytes(p))
	}
	ns := nodeList.NodeSet()

	addr := ecommon.Hex2Bytes(common.Replace0x(bscProof.Address))
	if !bytes.Equal(addr, contractAddr) {
		return nil, fmt.Errorf("verifyMerkleProof, contract address is error, proof address: %s, side chain address: %s", bscProof.Address, hex.EncodeToString(contractAddr))
	}
	acctKey := crypto.Keccak256(addr)

	//2. verify account proof
	acctVal, err := trie.VerifyProof(blockData.Root, acctKey, ns)
	if err != nil {
		return nil, fmt.Errorf("verifyMerkleProof, verify account proof error:%s", err)
	}

	nounce := new(big.Int)
	_, ok := nounce.SetString(common.Replace0x(bscProof.Nonce), 16)
	if !ok {
		return nil, fmt.Errorf("verifyMerkleProof, invalid format of nounce:%s", bscProof.Nonce)
	}

	balance := new(big.Int)
	_, ok = balance.SetString(common.Replace0x(bscProof.Balance), 16)
	if !ok {
		return nil, fmt.Errorf("verifyMerkleProof, invalid format of balance:%s", bscProof.Balance)
	}

	storageHash := ecommon.HexToHash(common.Replace0x(bscProof.StorageHash))
	codeHash := ecommon.HexToHash(common.Replace0x(bscProof.CodeHash))

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
	if len(bscProof.StorageProofs) != 1 {
		return nil, fmt.Errorf("verifyMerkleProof, invalid storage proof format")
	}

	sp := bscProof.StorageProofs[0]
	storageKey := crypto.Keccak256(ecommon.HexToHash(common.Replace0x(sp.Key)).Bytes())

	for _, prf := range sp.Proof {
		nodeList.Put(nil, ecommon.Hex2Bytes(common.Replace0x(prf)))
	}

	ns = nodeList.NodeSet()
	val, err := trie.VerifyProof(storageHash, storageKey, ns)
	if err != nil {
		return nil, fmt.Errorf("verifyMerkleProof, verify storage proof error:%s", err)
	}

	return val, nil
}

func checkProofResult(result, value []byte) bool {
	var tempBytes []byte
	err := rlp.DecodeBytes(result, &tempBytes)
	if err != nil {
		log.NewNopLogger().Error("checkProofResult, rlp.DecodeBytes error:%s\n", err.Error())
		return false
	}
	//
	var s []byte
	for i := len(tempBytes); i < 32; i++ {
		s = append(s, 0)
	}
	s = append(s, tempBytes...)
	hash := crypto.Keccak256(value)
	return bytes.Equal(s, hash)
}

// MakeDepositProposal ...
func (h *Handler) MakeDepositProposal(params common.EntranceParam) (*common.MakeTxParam, error) {
	// sourceChainID := params.SourceChainID
	height := params.Height
	proof := params.Proof
	// relayerAddress := params.RelayerAddress
	extra := params.Extra
	// headerOrCrossChainMsg := params.HeaderOrCrossChainMsg
	canonicalHeight := params.CanonicalHeight
	CCMCAddress := params.CCMCAddress
	blocksToWait := params.BlocksToWait
	headerWithSum := params.HeaderWithSum

	if canonicalHeight < height || canonicalHeight-height < uint32(blocksToWait-1) {
		return nil, fmt.Errorf("verifyFromTx, transaction is not confirmed, current height: %d, input height: %d", canonicalHeight, height)
	}

	bscProof := new(Proof)
	err := json.Unmarshal(proof, bscProof)
	if err != nil {
		return nil, fmt.Errorf("verifyFromTx, unmarshal proof error:%s", err)
	}

	if len(bscProof.StorageProofs) != 1 {
		return nil, fmt.Errorf("verifyFromTx, incorrect proof format")
	}

	proofResult, err := verifyMerkleProof(bscProof, headerWithSum.Header, CCMCAddress)
	if err != nil {
		return nil, fmt.Errorf("verifyFromTx, verifyMerkleProof error:%v", err)
	}

	if proofResult == nil {
		return nil, fmt.Errorf("verifyFromTx, verifyMerkleProof failed")
	}

	if !checkProofResult(proofResult, extra) {
		return nil, fmt.Errorf("verifyFromTx, verify proof value hash failed, proof result:%x, extra:%x", proofResult, extra)
	}

	data := types.NewZeroCopySource(extra)
	txParam := new(common.MakeTxParam)
	if err := txParam.Deserialization(data); err != nil {
		return nil, fmt.Errorf("verifyFromTx, deserialize merkleValue error:%s", err)
	}
	return txParam, nil
}

// SyncGenesisHeader ...
func (h *Handler) SyncGenesisHeader(params common.SyncGenesisHeaderParam) (*common.GenesisHeader, error) {
	var genesis common.GenesisHeader

	err := json.Unmarshal(params.GenesisHeader, &genesis)
	if err != nil {
		return &genesis, errors.New("bsc Handler SyncGenesisHeader, deserialize GenesisHeader err")
	}

	signersBytes := len(genesis.Header.Extra) - extraVanity - extraSeal
	if signersBytes == 0 || signersBytes%ecommon.AddressLength != 0 {
		return &genesis, errors.New("invalid signer list")
	}

	if len(genesis.PrevValidators) != 1 {
		return &genesis, errors.New("invalid PrevValidators")
	}
	if genesis.Header.Number.Cmp(genesis.PrevValidators[0].Height) <= 0 {
		return &genesis, errors.New("invalid height orders")
	}
	validators, err := ParseValidators(genesis.Header.Extra[extraVanity : extraVanity+signersBytes])
	if err != nil {
		return &genesis, err
	}
	genesis.PrevValidators = append([]common.HeightAndValidators{
		{Height: genesis.Header.Number, Validators: validators},
	}, genesis.PrevValidators...)

	return &genesis, nil
}

// ParseValidators ...
func ParseValidators(validatorsBytes []byte) ([]ecommon.Address, error) {
	if len(validatorsBytes)%ecommon.AddressLength != 0 {
		return nil, errors.New("invalid validators bytes")
	}
	n := len(validatorsBytes) / ecommon.AddressLength
	result := make([]ecommon.Address, n)
	for i := 0; i < n; i++ {
		address := make([]byte, ecommon.AddressLength)
		copy(address, validatorsBytes[i*ecommon.AddressLength:(i+1)*ecommon.AddressLength])
		result[i] = ecommon.BytesToAddress(address)
	}
	return result, nil
}

// SyncBlockHeader ...
func (h *Handler) SyncBlockHeader(params common.SyncBlockHeaderParam) (etypes.Header, error) {
	var header etypes.Header

	err := json.Unmarshal(params.Header, &header)
	if err != nil {
		return header, nil
	}

	var (
		inTurnHV common.HeightAndValidators
	)

	diffWithLastEpoch := big.NewInt(0).Sub(header.Number, params.Phv.Height).Int64()
	if diffWithLastEpoch <= int64(len(params.Pphv.Validators)/2) {
		// pphv is in effect
		inTurnHV = params.Pphv
	} else {
		// phv is in effect
		inTurnHV = params.Phv
	}

	if params.LastSeenHeight > 0 {
		limit := int64(len(inTurnHV.Validators) / 2)
		if header.Number.Int64() <= params.LastSeenHeight+limit {
			return header, nil
		}
	}

	indexInTurn := int(header.Number.Uint64()) % len(inTurnHV.Validators)
	if indexInTurn < 0 {
		return header, nil
	}
	valid := false
	for idx, v := range inTurnHV.Validators {
		if v == params.Signer {
			valid = true
			if indexInTurn == idx {
				if header.Difficulty.Cmp(diffInTurn) != 0 {
					break
				}
			} else {
				if header.Difficulty.Cmp(diffNoTurn) != 0 {
					break
				}
			}
		}
	}
	if !valid {
		return header, nil
	}

	return header, nil
}

// SyncCrossChainMsg ...
func (h *Handler) SyncCrossChainMsg(params common.SyncCrossChainMsgParam) error {
	return nil
}
