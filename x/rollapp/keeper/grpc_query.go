package keeper

import (
	"github.com/furychain/furya/x/rollapp/types"
)

var _ types.QueryServer = Keeper{}
