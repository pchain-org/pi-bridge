package blockchain

import (
	"fmt"

	"github.com/pchain-org/pi-bridge/tools/blockchain/bsc"
	"github.com/pchain-org/pi-bridge/tools/blockchain/common"
	"github.com/pchain-org/pi-bridge/tools/blockchain/eth"
	"github.com/pchain-org/pi-bridge/tools/blockchain/icp"
	"github.com/pchain-org/pi-bridge/tools/blockchain/plian"
)

var (
	ETH_ROUTER   = uint64(1)
	BSC_ROUTER   = uint64(2)
	PLIAN_ROUTER = uint64(3)
	ICP_ROUTER   = uint64(4)
)

func GetChainHandler(router uint64) (common.ChainHandler, error) {
	switch router {
	case ETH_ROUTER:
		return eth.NewHandler(), nil
	case BSC_ROUTER:
		return bsc.NewHandler(), nil
	case PLIAN_ROUTER:
		return plian.NewHandler(), nil
	case ICP_ROUTER:
		return icp.NewHandler(), nil
	default:
		return nil, fmt.Errorf("not a supported router:%d", router)
	}
}
