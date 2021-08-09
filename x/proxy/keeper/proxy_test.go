package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/pchain-org/pi-bridge/x/proxy/types"
)

func createNProxy(keeper *Keeper, ctx sdk.Context, n int) []types.Proxy {
	items := make([]types.Proxy, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetProxy(ctx, items[i])
	}
	return items
}

func TestProxyGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProxy(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProxy(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestProxyRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProxy(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProxy(ctx, item.Index)
		_, found := keeper.GetProxy(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestProxyGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProxy(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllProxy(ctx))
}
