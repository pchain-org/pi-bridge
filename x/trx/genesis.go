package trx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pchain-org/pi-bridge/x/trx/keeper"
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the trx
	for _, elem := range genState.TrxList {
		k.SetTrx(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all trx
	trxList := k.GetAllTrx(ctx)
	for _, elem := range trxList {
		elem := elem
		genesis.TrxList = append(genesis.TrxList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
