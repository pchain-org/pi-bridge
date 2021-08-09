package keeper

import (
	"github.com/pchain-org/pi-bridge/x/node/types"
)

var _ types.QueryServer = Keeper{}
