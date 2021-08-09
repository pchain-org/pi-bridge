package keeper

import (
	"github.com/pchain-org/pi-bridge/x/chain/types"
)

var _ types.QueryServer = Keeper{}
