package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/proxy/types"
)

// SetProxy set a specific proxy in the store from its index
func (k Keeper) SetProxy(ctx sdk.Context, proxy types.Proxy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProxyKey))
	b := k.cdc.MustMarshal(&proxy)
	store.Set(types.KeyPrefix(proxy.Index), b)
}

// GetProxy returns a proxy from its index
func (k Keeper) GetProxy(ctx sdk.Context, index string) (val types.Proxy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProxyKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.UnmarshalInterface(b, &val)
	return val, true
}

// RemoveProxy removes a proxy from the store
func (k Keeper) RemoveProxy(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProxyKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllProxy returns all proxy
func (k Keeper) GetAllProxy(ctx sdk.Context) (list []types.Proxy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProxyKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proxy
		k.cdc.UnmarshalInterface(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
