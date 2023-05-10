package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xblackfury/furya/testutil/keeper"
	"github.com/xblackfury/furya/testutil/nullify"
	"github.com/xblackfury/furya/x/rollapp/keeper"
	"github.com/xblackfury/furya/x/rollapp/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLatestFinalizedStateIndex(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StateInfoIndex {
	var items []types.StateInfoIndex
	stateInfoList := make([]types.StateInfo, n)
	for i := range stateInfoList {
		stateInfoList[i].StateInfoIndex.RollappId = strconv.Itoa(i)
		stateInfoList[i].StateInfoIndex.Index = uint64(i)
		stateInfoList[i].Status = types.STATE_STATUS_FINALIZED

		keeper.SetStateInfo(ctx, stateInfoList[i])
		keeper.SetLatestFinalizedStateIndex(ctx, stateInfoList[i].StateInfoIndex)

		items = append(items, stateInfoList[i].StateInfoIndex)
	}
	return items
}

func TestLatestFinalizedStateIndexGet(t *testing.T) {
	keeper, ctx := keepertest.RollappKeeper(t)
	items := createNLatestFinalizedStateIndex(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLatestFinalizedStateIndex(ctx,
			item.RollappId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLatestFinalizedStateIndexRemove(t *testing.T) {
	keeper, ctx := keepertest.RollappKeeper(t)
	items := createNLatestFinalizedStateIndex(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLatestFinalizedStateIndex(ctx,
			item.RollappId,
		)
		_, found := keeper.GetLatestFinalizedStateIndex(ctx,
			item.RollappId,
		)
		require.False(t, found)
	}
}

func TestLatestFinalizedStateIndexGetAll(t *testing.T) {
	keeper, ctx := keepertest.RollappKeeper(t)
	items := createNLatestFinalizedStateIndex(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLatestFinalizedStateIndex(ctx)),
	)
}
