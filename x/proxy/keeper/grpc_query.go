package keeper

import (
	"github.com/pchain-org/pi-bridge/x/proxy/types"
)

var _ types.QueryServer = Keeper{}
