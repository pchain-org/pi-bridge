package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/pchain-org/pi-bridge/x/block/types"
)

func createNBlock(keeper *Keeper, ctx sdk.Context, n int) []types.Block {
	items := make([]types.Block, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetBlock(ctx, items[i])
	}
	return items
}

func TestBlockGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBlock(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestBlockRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBlock(ctx, item.Index)
		_, found := keeper.GetBlock(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestBlockGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllBlock(ctx))
}
