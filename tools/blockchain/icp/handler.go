package icp

import (
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
)

// Handler ...
type Handler struct {
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

// MakeDepositProposal ...
func (h *Handler) MakeDepositProposal(params common.EntranceParam) (*common.MakeTxParam, error) {
	var txParam common.MakeTxParam
	return &txParam, nil
}

// SyncGenesisHeader ...
func (h *Handler) SyncGenesisHeader(params common.SyncGenesisHeaderParam) (*common.GenesisHeader, error) {
	var genesis common.GenesisHeader

	return &genesis, nil
}

// SyncBlockHeader ...
func (h *Handler) SyncBlockHeader(params common.SyncBlockHeaderParam) (etypes.Header, error) {
	var header etypes.Header

	return header, nil
}

// SyncCrossChainMsg ...
func (h *Handler) SyncCrossChainMsg(params common.SyncCrossChainMsgParam) error {
	return nil
}
