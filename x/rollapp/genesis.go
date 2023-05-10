package rollapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xblackfury/furya/x/rollapp/keeper"
	"github.com/xblackfury/furya/x/rollapp/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the rollapp
	for _, elem := range genState.RollappList {
		k.SetRollapp(ctx, elem)
	}
	// Set all the stateInfo
	for _, elem := range genState.StateInfoList {
		k.SetStateInfo(ctx, elem)
	}
	// Set all the latestStateInfoIndex
	for _, elem := range genState.LatestStateInfoIndexList {
		k.SetLatestStateInfoIndex(ctx, elem)
	}
	// Set all the latestFinalizedStateIndex
	for _, elem := range genState.LatestFinalizedStateIndexList {
		k.SetLatestFinalizedStateIndex(ctx, elem)
	}
	// Set all the blockHeightToFinalizationQueue
	for _, elem := range genState.BlockHeightToFinalizationQueueList {
		k.SetBlockHeightToFinalizationQueue(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RollappList = k.GetAllRollapp(ctx)
	genesis.StateInfoList = k.GetAllStateInfo(ctx)
	genesis.LatestStateInfoIndexList = k.GetAllLatestStateInfoIndex(ctx)
	genesis.LatestFinalizedStateIndexList = k.GetAllLatestFinalizedStateIndex(ctx)
	genesis.BlockHeightToFinalizationQueueList = k.GetAllBlockHeightToFinalizationQueue(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
