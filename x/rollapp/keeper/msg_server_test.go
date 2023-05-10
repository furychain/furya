package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xblackfury/furya/testutil/keeper"
	"github.com/xblackfury/furya/x/rollapp/keeper"
	"github.com/xblackfury/furya/x/rollapp/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RollappKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
