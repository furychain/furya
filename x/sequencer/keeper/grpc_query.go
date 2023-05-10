package keeper

import (
	"github.com/xblackfury/furya/x/sequencer/types"
)

var _ types.QueryServer = Keeper{}
