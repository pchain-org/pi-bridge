package keeper

import (
	"github.com/pchain-org/pi-bridge/x/block/types"
)

var _ types.QueryServer = Keeper{}
