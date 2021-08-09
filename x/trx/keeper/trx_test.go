package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/pchain-org/pi-bridge/x/trx/types"
)

func createNTrx(keeper *Keeper, ctx sdk.Context, n int) []types.Trx {
	items := make([]types.Trx, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetTrx(ctx, items[i])
	}
	return items
}

func TestTrxGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNTrx(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTrx(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestTrxRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNTrx(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTrx(ctx, item.Index)
		_, found := keeper.GetTrx(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestTrxGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNTrx(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllTrx(ctx))
}
