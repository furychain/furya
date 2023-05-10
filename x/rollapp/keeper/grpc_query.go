package keeper

import (
	"github.com/xblackfury/furya/x/rollapp/types"
)

var _ types.QueryServer = Keeper{}
