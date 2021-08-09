package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/chain/types"
)

// SetChain set a specific chain in the store from its index
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	b := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(types.KeyPrefix(chain.Index), b)
}

// GetChain returns a chain from its index
func (k Keeper) GetChain(ctx sdk.Context, index string) (val types.Chain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveChain removes a chain from the store
func (k Keeper) RemoveChain(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllChain returns all chain
func (k Keeper) GetAllChain(ctx sdk.Context) (list []types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Chain
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
