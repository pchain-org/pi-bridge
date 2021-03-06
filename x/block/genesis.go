package block

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/block/keeper"
	"github.com/pchain-org/pi-bridge/x/block/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the block
	for _, elem := range genState.BlockList {
		k.SetBlock(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all block
	blockList := k.GetAllBlock(ctx)
	for _, elem := range blockList {
		elem := elem
		genesis.BlockList = append(genesis.BlockList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
