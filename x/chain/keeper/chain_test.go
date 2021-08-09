package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/pchain-org/pi-bridge/x/chain/types"
)

func createNChain(keeper *Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func TestChainGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChain(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestChainRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.Index)
		_, found := keeper.GetChain(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestChainGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllChain(ctx))
}
