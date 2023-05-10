package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xblackfury/furya/testutil/keeper"
	"github.com/xblackfury/furya/x/sequencer/keeper"
	"github.com/xblackfury/furya/x/sequencer/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SequencerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
