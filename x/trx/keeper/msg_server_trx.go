package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

func (k msgServer) CreateTrx(goCtx context.Context, msg *types.MsgCreateTrx) (*types.MsgCreateTrxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTrx(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.Index))
	}

	var trx = types.Trx{
		Index:          msg.Index,
		Creator:        msg.Creator,
		ChainID:        msg.ChainID,
		Address:        msg.Address,
		CrossChainMsgs: msg.CrossChainMsgs,
	}

	k.SetTrx(
		ctx,
		trx,
	)
	return &types.MsgCreateTrxResponse{}, nil
}

func (k msgServer) UpdateTrx(goCtx context.Context, msg *types.MsgUpdateTrx) (*types.MsgUpdateTrxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTrx(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var trx = types.Trx{
		Index:          msg.Index,
		Creator:        msg.Creator,
		ChainID:        msg.ChainID,
		Address:        msg.Address,
		CrossChainMsgs: msg.CrossChainMsgs,
	}

	k.SetTrx(ctx, trx)

	return &types.MsgUpdateTrxResponse{}, nil
}

func (k msgServer) DeleteTrx(goCtx context.Context, msg *types.MsgDeleteTrx) (*types.MsgDeleteTrxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTrx(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTrx(ctx, msg.Index)

	return &types.MsgDeleteTrxResponse{}, nil
}
