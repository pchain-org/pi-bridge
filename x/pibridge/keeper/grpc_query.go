package keeper

import (
	"github.com/pchain-org/pi-bridge/x/pibridge/types"
)

var _ types.QueryServer = Keeper{}
