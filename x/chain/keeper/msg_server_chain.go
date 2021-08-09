package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pchain-org/pi-bridge/x/chain/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetChain(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.Index))
	}

	var chain = types.Chain{
		Index:                 msg.Index,
		Creator:               msg.Creator,
		SourceChainID:         msg.SourceChainID,
		Height:                msg.Height,
		Proof:                 msg.Proof,
		ProxyAddress:          msg.ProxyAddress,
		Extra:                 msg.Extra,
		HeaderOrCrossChainMsg: msg.HeaderOrCrossChainMsg,
	}

	k.SetChain(
		ctx,
		chain,
	)
	return &types.MsgCreateChainResponse{}, nil
}

func (k msgServer) UpdateChain(goCtx context.Context, msg *types.MsgUpdateChain) (*types.MsgUpdateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetChain(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var chain = types.Chain{
		Index:                 msg.Index,
		Creator:               msg.Creator,
		SourceChainID:         msg.SourceChainID,
		Height:                msg.Height,
		Proof:                 msg.Proof,
		ProxyAddress:          msg.ProxyAddress,
		Extra:                 msg.Extra,
		HeaderOrCrossChainMsg: msg.HeaderOrCrossChainMsg,
	}

	k.SetChain(ctx, chain)

	return &types.MsgUpdateChainResponse{}, nil
}

func (k msgServer) DeleteChain(goCtx context.Context, msg *types.MsgDeleteChain) (*types.MsgDeleteChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetChain(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveChain(ctx, msg.Index)

	return &types.MsgDeleteChainResponse{}, nil
}
