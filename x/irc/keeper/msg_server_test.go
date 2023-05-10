package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xblackfury/furya/testutil/keeper"
	"github.com/xblackfury/furya/x/irc/keeper"
	"github.com/xblackfury/furya/x/irc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, _, ctx := keepertest.IRCKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
