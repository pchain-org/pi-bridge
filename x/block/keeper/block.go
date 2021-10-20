package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/block/types"
)

// SetBlock set a specific block in the store from its index
func (k Keeper) SetBlock(ctx sdk.Context, block types.Block) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	b := k.cdc.MustMarshal(&block)
	store.Set(types.KeyPrefix(block.Index), b)
}

// GetBlock returns a block from its index
func (k Keeper) GetBlock(ctx sdk.Context, index string) (val types.Block, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.UnmarshalInterface(b, &val)
	return val, true
}

// RemoveBlock removes a block from the store
func (k Keeper) RemoveBlock(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllBlock returns all block
func (k Keeper) GetAllBlock(ctx sdk.Context) (list []types.Block) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Block
		k.cdc.UnmarshalInterface(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
