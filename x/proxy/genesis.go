package proxy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/proxy/keeper"
	"github.com/pchain-org/pi-bridge/x/proxy/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the proxy
	for _, elem := range genState.ProxyList {
		k.SetProxy(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all proxy
	proxyList := k.GetAllProxy(ctx)
	for _, elem := range proxyList {
		elem := elem
		genesis.ProxyList = append(genesis.ProxyList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
