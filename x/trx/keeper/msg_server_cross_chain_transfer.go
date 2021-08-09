package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

func (k msgServer) CrossChainTransfer(goCtx context.Context, msg *types.MsgCrossChainTransfer) (*types.MsgCrossChainTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCrossChainTransferResponse{}, nil
}
