package common

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ecommon "github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pchain-org/pi-bridge/tools/types"
)

type MakeTxParam struct {
	TxHash              []byte
	CrossChainID        []byte
	FromContractAddress []byte
	ToChainID           uint64
	ToContractAddress   []byte
	Method              string
	Args                []byte
}

func (this *MakeTxParam) Serialization(sink *types.ZeroCopySink) {
	sink.WriteVarBytes(this.TxHash)
	sink.WriteVarBytes(this.CrossChainID)
	sink.WriteVarBytes(this.FromContractAddress)
	sink.WriteUint64(this.ToChainID)
	sink.WriteVarBytes(this.ToContractAddress)
	sink.WriteVarBytes([]byte(this.Method))
	sink.WriteVarBytes(this.Args)
}

func (this *MakeTxParam) Deserialization(source *types.ZeroCopySource) error {
	txHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize txHash error")
	}
	crossChainID, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize crossChainID error")
	}
	fromContractAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize fromContractAddress error")
	}
	toChainID, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize toChainID error")
	}
	toContractAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize toContractAddress error")
	}
	method, eof := source.NextString()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize method error")
	}
	args, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize args error")
	}

	this.TxHash = txHash
	this.CrossChainID = crossChainID
	this.FromContractAddress = fromContractAddress
	this.ToChainID = toChainID
	this.ToContractAddress = toContractAddress
	this.Method = method
	this.Args = args
	return nil
}

// GenesisHeader ...
type GenesisHeader struct {
	Header         etypes.Header
	PrevValidators []HeightAndValidators
}

// HeightAndValidators ...
type HeightAndValidators struct {
	Height     *big.Int
	Validators []ecommon.Address
	Hash       *ecommon.Hash
}

// HeaderWithDifficultySum ...
type HeaderWithDifficultySum struct {
	Header          *etypes.Header `json:"header"`
	DifficultySum   *big.Int       `json:"difficultySum"`
	EpochParentHash *ecommon.Hash  `json:"epochParentHash"`
}

type EntranceParam struct {
	SourceChainID         uint64 `json:"sourceChainId"`
	Height                uint32 `json:"height"`
	Proof                 []byte `json:"proof"`
	RelayerAddress        []byte `json:"relayerAddress"`
	Extra                 []byte `json:"extra"`
	HeaderOrCrossChainMsg []byte `json:"headerOrCrossChainMsg"`
	//bsc
	CanonicalHeight uint32                   `json:"canonicalHeight"`
	CCMCAddress     []byte                   `json:"cCMCAddress"`
	BlocksToWait    uint64                   `json:"blocksToWait"`
	HeaderWithSum   *HeaderWithDifficultySum `json:"headerWithSum"`
	//eth
	ETHHeader *Header `json:"ethHeader"`
}

type SyncGenesisHeaderParam struct {
	ChainID       uint64 `json:"chainId"`
	GenesisHeader []byte `json:"genesisHeader"`
}
type SyncBlockHeaderParam struct {
	ChainID             uint64              `json:"chainId"`
	GenesisHeader       GenesisHeader       `json:"genesisHeader"`
	Address             sdk.Address         `json:"address"`
	Header              []byte              `json:"header"`
	Phv                 HeightAndValidators `json:"phv"`
	Pphv                HeightAndValidators `json:"pphv"`
	LastSeenHeight      int64               `json:"lastSeenHeight"`
	Signer              ecommon.Address     `json:"signer"`
	ExtraInfo           ExtraInfo           `json:"extraInfo"`
	ParentHeader        etypes.Header       `json:"parentHeader"`
	ParentDifficultySum big.Int             `json:"parentDifficultySum"`
}
type SyncCrossChainMsgParam struct {
	ChainID        uint64      `json:"chainId"`
	Address        sdk.Address `json:"address"`
	CrossChainMsgs [][]byte    `json:"crossChainMsgs"`
}

// ExtraInfo ...
type ExtraInfo struct {
	ChainID *big.Int // for bsc
}

// Context ...
type Context struct {
	ExtraInfo ExtraInfo
	ChainID   uint64
}

type ChainHandler interface {
	MakeDepositProposal(EntranceParam) (*MakeTxParam, error)
	SyncGenesisHeader(SyncGenesisHeaderParam) (*GenesisHeader, error)
	SyncBlockHeader(SyncBlockHeaderParam) (etypes.Header, error)
	SyncCrossChainMsg(SyncCrossChainMsgParam) error
}
