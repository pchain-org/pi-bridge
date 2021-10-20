package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

// SetTrx set a specific trx in the store from its index
func (k Keeper) SetTrx(ctx sdk.Context, trx types.Trx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrxKey))
	b := k.cdc.MustMarshal(&trx)
	store.Set(types.KeyPrefix(trx.Index), b)
}

// GetTrx returns a trx from its index
func (k Keeper) GetTrx(ctx sdk.Context, index string) (val types.Trx, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrxKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.UnmarshalInterface(b, &val)
	return val, true
}

// RemoveTrx removes a trx from the store
func (k Keeper) RemoveTrx(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrxKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllTrx returns all trx
func (k Keeper) GetAllTrx(ctx sdk.Context) (list []types.Trx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrxKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Trx
		k.cdc.UnmarshalInterface(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
