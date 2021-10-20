package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/pchain-org/pi-bridge/x/proxy/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProxyAll(c context.Context, req *types.QueryAllProxyRequest) (*types.QueryAllProxyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proxys []*types.Proxy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proxyStore := prefix.NewStore(store, types.KeyPrefix(types.ProxyKey))

	pageRes, err := query.Paginate(proxyStore, req.Pagination, func(key []byte, value []byte) error {
		var proxy types.Proxy
		if err := k.cdc.UnmarshalInterface(value, &proxy); err != nil {
			return err
		}

		proxys = append(proxys, &proxy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProxyResponse{Proxy: proxys, Pagination: pageRes}, nil
}

func (k Keeper) Proxy(c context.Context, req *types.QueryGetProxyRequest) (*types.QueryGetProxyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProxy(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetProxyResponse{Proxy: &val}, nil
}
