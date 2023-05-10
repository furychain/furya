package keeper

import (
	"github.com/furychain/furya/x/irc/types"
)

var _ types.QueryServer = Keeper{}
