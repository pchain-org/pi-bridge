package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pchain-org/pi-bridge/x/node/types"
)

func (k msgServer) CreateNode(goCtx context.Context, msg *types.MsgCreateNode) (*types.MsgCreateNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetNode(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.Index))
	}

	var node = types.Node{
		Index:      msg.Index,
		Creator:    msg.Creator,
		PeerPubkey: msg.PeerPubkey,
		Address:    msg.Address,
	}

	k.SetNode(
		ctx,
		node,
	)
	return &types.MsgCreateNodeResponse{}, nil
}

func (k msgServer) UpdateNode(goCtx context.Context, msg *types.MsgUpdateNode) (*types.MsgUpdateNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNode(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var node = types.Node{
		Index:      msg.Index,
		Creator:    msg.Creator,
		PeerPubkey: msg.PeerPubkey,
		Address:    msg.Address,
	}

	k.SetNode(ctx, node)

	return &types.MsgUpdateNodeResponse{}, nil
}

func (k msgServer) DeleteNode(goCtx context.Context, msg *types.MsgDeleteNode) (*types.MsgDeleteNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNode(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveNode(ctx, msg.Index)

	return &types.MsgDeleteNodeResponse{}, nil
}
