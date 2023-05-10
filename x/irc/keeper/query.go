package keeper

import (
	"github.com/xblackfury/furya/x/irc/types"
)

var _ types.QueryServer = Keeper{}
