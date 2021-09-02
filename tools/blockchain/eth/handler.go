package eth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/consensus"
	etypes "github.com/ethereum/go-ethereum/core/types"
	eparams "github.com/ethereum/go-ethereum/params"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/types"
)

// Handler ...
type Handler struct {
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) MakeDepositProposal(params common.EntranceParam) (*common.MakeTxParam, error) {

	blockData := params.ETHHeader

	ethProof := new(common.ETHProof)
	err := json.Unmarshal(params.Proof, ethProof)
	if err != nil {
		return nil, fmt.Errorf("VerifyFromEthProof, unmarshal proof error:%s", err)
	}

	if len(ethProof.StorageProofs) != 1 {
		return nil, fmt.Errorf("VerifyFromEthProof, incorrect proof format")
	}

	//todo 1. verify the proof with header
	//determine where the k and v from
	proofResult, err := common.VerifyMerkleProof(ethProof, blockData, params.CCMCAddress)
	if err != nil {
		return nil, fmt.Errorf("VerifyFromEthProof, verifyMerkleProof error:%v", err)
	}
	if proofResult == nil {
		return nil, fmt.Errorf("VerifyFromEthProof, verifyMerkleProof failed!\n")
	}

	if !common.CheckProofResult(proofResult, params.Extra) {
		return nil, fmt.Errorf("VerifyFromEthProof, verify proof value hash failed, proof result:%x, extra:%x", proofResult, params.Extra)
	}

	data := types.NewZeroCopySource(params.Extra)
	txParam := new(common.MakeTxParam)
	if err := txParam.Deserialization(data); err != nil {
		return nil, fmt.Errorf("VerifyFromEthProof, deserialize merkleValue error:%s", err)
	}
	return txParam, nil
}

// SyncGenesisHeader ...
func (h *Handler) SyncGenesisHeader(params common.SyncGenesisHeaderParam) (*common.GenesisHeader, error) {
	var genesis common.GenesisHeader
	var header etypes.Header

	err := json.Unmarshal(params.GenesisHeader, &header)
	if err != nil {
		return &genesis, errors.New("ETH Handler SyncGenesisHeader, deserialize GenesisHeader err")
	}

	genesis.Header = header
	return &genesis, nil
}

// SyncBlockHeader ...
func (this *Handler) SyncBlockHeader(params common.SyncBlockHeaderParam) (etypes.Header, error) {

	var header etypes.Header
	err := json.Unmarshal(params.Header, &header)
	if err != nil {
		return header, fmt.Errorf("SyncBlockHeader, deserialize header err: %v", err)
	}
	// headerHash := header.Hash()
	// get pre header
	parentHeader := params.ParentHeader
	// parentDifficultySum := params.ParentDifficultySum
	parentHeaderHash := parentHeader.Hash()
	/**
	this code source refer to https://github.com/ethereum/go-ethereum/blob/master/consensus/ethash/consensus.go
	verify header need to verify:
	1. parent hash
	2. extra size
	3. current time
	*/
	//verify whether parent hash validity
	if !bytes.Equal(parentHeaderHash.Bytes(), header.ParentHash.Bytes()) {
		return header, fmt.Errorf("SyncBlockHeader, parent header is not right. Header: %s", string(params.Header))
	}
	//verify whether extra size validity
	if uint64(len(header.Extra)) > eparams.MaximumExtraDataSize {
		return header, fmt.Errorf("SyncBlockHeader, SyncBlockHeader extra-data too long: %d > %d, header: %s", len(header.Extra), eparams.MaximumExtraDataSize, string(params.Header))
	}
	//verify current time validity
	if header.Time > uint64(time.Now().Add(common.AllowedFutureBlockTime).Unix()) {
		return header, fmt.Errorf("SyncBlockHeader,  verify header time error:%s, checktime: %d, header: %s", consensus.ErrFutureBlock, time.Now().Add(common.AllowedFutureBlockTime).Unix(), string(params.Header))
	}
	//verify whether current header time and prevent header time validity
	if header.Time <= parentHeader.Time {
		return header, fmt.Errorf("SyncBlockHeader, verify header time fail. Header: %s", string(params.Header))
	}
	// Verify that the gas limit is <= 2^63-1
	cap := uint64(0x7fffffffffffffff)
	if header.GasLimit > cap {
		return header, fmt.Errorf("SyncBlockHeader, invalid gasLimit: have %v, max %v, header: %s", header.GasLimit, cap, string(params.Header))
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return header, fmt.Errorf("SyncBlockHeader, invalid gasUsed: have %d, gasLimit %d, header: %s", header.GasUsed, header.GasLimit, string(params.Header))
	}
	cheader := common.ConvertEHeaderToHeader(header)
	cParentHeader := common.ConvertEHeaderToHeader(parentHeader)
	if common.IsLondon(&cheader) {
		err = common.VerifyEip1559Header(&cParentHeader, &cheader)
	} else {
		err = common.VerifyGaslimit(parentHeader.GasLimit, cheader.GasLimit)
	}
	if err != nil {
		return header, fmt.Errorf("SyncBlockHeader, err:%v", err)
	}

	//verify difficulty
	var expected *big.Int
	if common.IsLondon(&cheader) {
		expected = common.MakeDifficultyCalculator(big.NewInt(9700000))(cheader.Time, &cParentHeader)
	} else {
		expected = common.DifficultyCalculator(new(big.Int).SetUint64(cheader.Time), &cParentHeader)
	}
	if expected.Cmp(header.Difficulty) != 0 {
		return header, fmt.Errorf("SyncBlockHeader, invalid difficulty: have %v, want %v, header: %s", header.Difficulty, expected, string(params.Header))
	}

	return header, nil

}

// SyncCrossChainMsg ...
func (this *Handler) SyncCrossChainMsg(params common.SyncCrossChainMsgParam) error {
	return nil
}
