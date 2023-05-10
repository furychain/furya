package keeper_test

import (
	"testing"

	testkeeper "github.com/xblackfury/furya/testutil/keeper"
	"github.com/xblackfury/furya/x/sequencer/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SequencerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
