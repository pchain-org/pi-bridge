package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pchain-org/pi-bridge/x/proxy/types"
)

func (k msgServer) CreateProxy(goCtx context.Context, msg *types.MsgCreateProxy) (*types.MsgCreateProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetProxy(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.Index))
	}

	var proxy = types.Proxy{
		Index:       msg.Index,
		Creator:     msg.Creator,
		AddressList: msg.AddressList,
		Address:     msg.Address,
	}

	k.SetProxy(
		ctx,
		proxy,
	)
	return &types.MsgCreateProxyResponse{}, nil
}

func (k msgServer) UpdateProxy(goCtx context.Context, msg *types.MsgUpdateProxy) (*types.MsgUpdateProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProxy(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var proxy = types.Proxy{
		Index:       msg.Index,
		Creator:     msg.Creator,
		AddressList: msg.AddressList,
		Address:     msg.Address,
	}

	k.SetProxy(ctx, proxy)

	return &types.MsgUpdateProxyResponse{}, nil
}

func (k msgServer) DeleteProxy(goCtx context.Context, msg *types.MsgDeleteProxy) (*types.MsgDeleteProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProxy(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProxy(ctx, msg.Index)

	return &types.MsgDeleteProxyResponse{}, nil
}
