package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/node/types"
)

// SetNode set a specific node in the store from its index
func (k Keeper) SetNode(ctx sdk.Context, node types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKey))
	b := k.cdc.MustMarshalBinaryBare(&node)
	store.Set(types.KeyPrefix(node.Index), b)
}

// GetNode returns a node from its index
func (k Keeper) GetNode(ctx sdk.Context, index string) (val types.Node, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveNode removes a node from the store
func (k Keeper) RemoveNode(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllNode returns all node
func (k Keeper) GetAllNode(ctx sdk.Context) (list []types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Node
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
