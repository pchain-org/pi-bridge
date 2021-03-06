package chain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/chain/keeper"
	"github.com/pchain-org/pi-bridge/x/chain/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the chain
	for _, elem := range genState.ChainList {
		k.SetChain(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all chain
	chainList := k.GetAllChain(ctx)
	for _, elem := range chainList {
		elem := elem
		genesis.ChainList = append(genesis.ChainList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
