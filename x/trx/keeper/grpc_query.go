package keeper

import (
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

var _ types.QueryServer = Keeper{}
