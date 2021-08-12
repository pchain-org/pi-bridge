package protocol

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	blockmodulekeeper "github.com/pchain-org/pi-bridge/x/block/keeper"
)

type BSCServer struct {
}
type Definty struct {
}

type ETHServer struct {
}

type PChainServer struct {
}
type Config struct {
}
type protocol struct {
	bsc           *BSCServer
	icp           *Definty
	eth           *ETHServer
	pi            *PChainServer
	ctx           sdk.Context
	config        *Config
	accountKeeper authkeeper.AccountKeeper
	blockKeeper   blockmodulekeeper.Keeper
}

func MonitorETHChain(ctx sdk.Context, config *Config, accountKeeper authkeeper.AccountKeeper,
	blockKeeper blockmodulekeeper.Keeper) {

	fmt.Println("***********************************************")
	fmt.Println("***********************************************")
	fmt.Println("***********************************************")
	fmt.Println("MonitorETHChain")
	fmt.Println("ctx.header.Height")
	fmt.Println(ctx.BlockHeight())
	height := strconv.FormatInt(ctx.BlockHeight(), 10)
	block, flag := blockKeeper.GetBlock(ctx, height)
	fmt.Println(block)
	fmt.Println(flag)
	
}

func MonitorETHCrossTX(ctx sdk.Context, config *Config, accountKeeper authkeeper.AccountKeeper,
	blockKeeper blockmodulekeeper.Keeper) {

	fmt.Println("#############################################")
	fmt.Println("#############################################")
	fmt.Println("#############################################")
	fmt.Println("MonitorETHCrossTX")
	fmt.Println("ctx")
	fmt.Println(ctx)
	fmt.Println("accountKeeper")
	fmt.Println("cosmos1gyplaw49rstrxc2q9xhukml8vjkdvparz4r6y6")
	add1, _ := sdk.AccAddressFromBech32("cosmos1gyplaw49rstrxc2q9xhukml8vjkdvparz4r6y6")
	add2 := accountKeeper.GetAccount(ctx, add1)
	fmt.Println(add2.GetPubKey())
	fmt.Println(add2.GetAccountNumber())
	fmt.Println(add2.GetSequence())
}
func (p *protocol) Start() {
	go MonitorETHChain(p.ctx, p.config, p.accountKeeper, p.blockKeeper)
	go MonitorETHCrossTX(p.ctx, p.config, p.accountKeeper, p.blockKeeper)
}

type ProtocolInstance interface {
	Start()
}

// func MakeProtocol(ctx sdk.Context, config *Config, accountKeeper authkeeper.AccountKeeper,
func MakeProtocol(ctx sdk.Context, accountKeeper authkeeper.AccountKeeper,
	blockKeeper blockmodulekeeper.Keeper,
) ProtocolInstance {
	p := &protocol{
		bsc:           &BSCServer{},
		icp:           &Definty{},
		eth:           &ETHServer{},
		pi:            &PChainServer{},
		ctx:           ctx,
		config:        &Config{},
		accountKeeper: accountKeeper,
		blockKeeper:   blockKeeper,
	}

	return p
}
