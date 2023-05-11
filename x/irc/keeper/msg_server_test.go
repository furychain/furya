package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/furychain/furya/testutil/keeper"
	"github.com/furychain/furya/x/irc/keeper"
	"github.com/furychain/furya/x/irc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, _, ctx := keepertest.IRCKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
