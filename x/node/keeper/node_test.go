package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/pchain-org/pi-bridge/x/node/types"
)

func createNNode(keeper *Keeper, ctx sdk.Context, n int) []types.Node {
	items := make([]types.Node, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetNode(ctx, items[i])
	}
	return items
}

func TestNodeGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNode(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNode(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestNodeRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNode(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNode(ctx, item.Index)
		_, found := keeper.GetNode(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestNodeGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNode(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllNode(ctx))
}
