package rollapp_test

import (
	"testing"

	keepertest "github.com/furychain/furya/testutil/keeper"
	"github.com/furychain/furya/testutil/nullify"
	"github.com/furychain/furya/x/rollapp"
	"github.com/furychain/furya/x/rollapp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RollappList: []types.Rollapp{
			{
				RollappId: "0",
			},
			{
				RollappId: "1",
			},
		},
		StateInfoList: []types.StateInfo{
			{
				StateInfoIndex: types.StateInfoIndex{
					RollappId: "0",
					Index:     0,
				},
			},
			{
				StateInfoIndex: types.StateInfoIndex{
					RollappId: "1",
					Index:     1,
				},
			},
		},
		LatestStateInfoIndexList: []types.StateInfoIndex{
			{
				RollappId: "0",
			},
			{
				RollappId: "1",
			},
		},
		BlockHeightToFinalizationQueueList: []types.BlockHeightToFinalizationQueue{
			{
				FinalizationHeight: 0,
			},
			{
				FinalizationHeight: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RollappKeeper(t)
	rollapp.InitGenesis(ctx, *k, genesisState)
	got := rollapp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RollappList, got.RollappList)
	require.ElementsMatch(t, genesisState.StateInfoList, got.StateInfoList)
	require.ElementsMatch(t, genesisState.LatestStateInfoIndexList, got.LatestStateInfoIndexList)
	require.ElementsMatch(t, genesisState.BlockHeightToFinalizationQueueList, got.BlockHeightToFinalizationQueueList)
	// this line is used by starport scaffolding # genesis/test/assert
}
