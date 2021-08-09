package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pchain-org/pi-bridge/x/block/types"
)

func (k msgServer) CreateBlock(goCtx context.Context, msg *types.MsgCreateBlock) (*types.MsgCreateBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetBlock(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.Index))
	}

	var block = types.Block{
		Index:   msg.Index,
		Creator: msg.Creator,
		ChainID: msg.ChainID,
		Address: msg.Address,
		Headers: msg.Headers,
	}

	k.SetBlock(
		ctx,
		block,
	)
	return &types.MsgCreateBlockResponse{}, nil
}

func (k msgServer) UpdateBlock(goCtx context.Context, msg *types.MsgUpdateBlock) (*types.MsgUpdateBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBlock(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var block = types.Block{
		Index:   msg.Index,
		Creator: msg.Creator,
		ChainID: msg.ChainID,
		Address: msg.Address,
		Headers: msg.Headers,
	}

	k.SetBlock(ctx, block)

	return &types.MsgUpdateBlockResponse{}, nil
}

func (k msgServer) DeleteBlock(goCtx context.Context, msg *types.MsgDeleteBlock) (*types.MsgDeleteBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBlock(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBlock(ctx, msg.Index)

	return &types.MsgDeleteBlockResponse{}, nil
}
