package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/pchain-org/pi-bridge/x/block/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BlockAll(c context.Context, req *types.QueryAllBlockRequest) (*types.QueryAllBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blocks []*types.Block
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	blockStore := prefix.NewStore(store, types.KeyPrefix(types.BlockKey))

	pageRes, err := query.Paginate(blockStore, req.Pagination, func(key []byte, value []byte) error {
		var block types.Block
		if err := k.cdc.UnmarshalInterface(value, &block); err != nil {
			return err
		}

		blocks = append(blocks, &block)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlockResponse{Block: blocks, Pagination: pageRes}, nil
}

func (k Keeper) Block(c context.Context, req *types.QueryGetBlockRequest) (*types.QueryGetBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBlock(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetBlockResponse{Block: &val}, nil
}
