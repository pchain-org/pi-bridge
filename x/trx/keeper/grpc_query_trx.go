package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/pchain-org/pi-bridge/x/trx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TrxAll(c context.Context, req *types.QueryAllTrxRequest) (*types.QueryAllTrxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var trxs []*types.Trx
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	trxStore := prefix.NewStore(store, types.KeyPrefix(types.TrxKey))

	pageRes, err := query.Paginate(trxStore, req.Pagination, func(key []byte, value []byte) error {
		var trx types.Trx
		if err := k.cdc.UnmarshalBinaryBare(value, &trx); err != nil {
			return err
		}

		trxs = append(trxs, &trx)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTrxResponse{Trx: trxs, Pagination: pageRes}, nil
}

func (k Keeper) Trx(c context.Context, req *types.QueryGetTrxRequest) (*types.QueryGetTrxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTrx(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetTrxResponse{Trx: &val}, nil
}
