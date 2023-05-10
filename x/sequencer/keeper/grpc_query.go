package keeper

import (
	"github.com/furychain/furya/x/sequencer/types"
)

var _ types.QueryServer = Keeper{}
