package protocol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	ethTools "github.com/pchain-org/pi-bridge/tools/blockchain/eth"
	"github.com/pchain-org/pi-bridge/tools/config"
	blockmodulekeeper "github.com/pchain-org/pi-bridge/x/block/keeper"
	chainmodulekeeper "github.com/pchain-org/pi-bridge/x/chain/keeper"
	nodemodulekeeper "github.com/pchain-org/pi-bridge/x/node/keeper"
	trxmodulekeeper "github.com/pchain-org/pi-bridge/x/trx/keeper"
	"github.com/tendermint/tendermint/crypto/merkle"
)

const (
	ChainETH = "ethereum"
	ChainBSC = "bsc"
	ChainPI  = "pchian"
	ChainICP = "definty"
)

///111111
type Contact struct {
	rpc string
}
type Response struct {
	Ret     int64           `json:"ret"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

func NewContact(rpc string) Contact {
	rpc = strings.TrimSuffix(rpc, "/")
	return Contact{rpc: rpc}
}

type QueryRes struct {
	Value  interface{}   `json:"value"`
	Height int64         `json:"height,omitempty"`
	Proof  *merkle.Proof `json:"proof,omitempty"`
}

// GetRequest defines a wrapper around an HTTP GET request with a provided URL.
// An error is returned if the request or reading the body fails.
func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url) // nolint:gosec
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = res.Body.Close(); err != nil {
		return nil, err
	}

	return body, nil
}

// PostRequest defines a wrapper around an HTTP POST request with a provided URL and data.
// An error is returned if the request or reading the body fails.
func PostRequest(url string, data url.Values) ([]byte, error) {
	res, err := http.PostForm(url, data) // nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("error while sending post request: %w", err)
	}

	bz, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err = res.Body.Close(); err != nil {
		return nil, err
	}

	return bz, nil
}
func (c Contact) Get(url string) ([]byte, error) {
	res, err := GetRequest(c.rpc + url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c Contact) Post(url string, data url.Values) ([]byte, error) {
	res, err := PostRequest(c.rpc+url, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c Contact) GetBalance(address string) sdk.Coins {
	data := url.Values{}
	data.Add("address", address)
	balanceRes, err := c.Post("/bank/balance", data)
	fmt.Println("balanceRes")
	fmt.Println(balanceRes)
	fmt.Println("err")
	fmt.Println(err)
	var res Response
	var queryRes QueryRes
	json.Unmarshal(balanceRes, &res)
	json.Unmarshal(res.Data, &queryRes)
	fmt.Println("res")
	fmt.Println(res)
	fmt.Println("queryRes")
	fmt.Println(queryRes)
	var balance sdk.Coins
	if len(res.Message) > 0 {
		balanceList := queryRes.Value.(map[string]interface{})["balance_list"].([]interface{})
		for _, v := range balanceList {
			vMap := v.(map[string]interface{})
			amt := vMap["amount"].(string)
			x, _ := new(big.Int).SetString(amt, 10)
			balance = append(balance, sdk.NewCoin(vMap["denom"].(string), sdk.NewIntFromBigInt(x)))
		}
	}

	return balance
}

///protocol
type protocol struct {
	ctx           sdk.Context
	cdc           codec.Codec
	config        *config.ServiceConfig
	accountKeeper authkeeper.AccountKeeper
	bankkeeper    bankkeeper.Keeper
	blockKeeper   blockmodulekeeper.Keeper
	chainKeeper   chainmodulekeeper.Keeper
	nodeKeeper    nodemodulekeeper.Keeper
	trxKeeper     trxmodulekeeper.Keeper
}

func (p *protocol) MonitorChain(chain string) {
	fmt.Println("MonitorChain")
	switch chain {
	case ChainETH:
		// chain, _ := p.chainKeeper.GetChain(p.ctx, chain)
		// ethManager, err := ethTools.NewEthManager(p.ctx, p.cdc, p.config, p.accountKeeper)
		// ethManager.SyncHeadersAndEvents(uint64(chain.Height))
		add1, _ := sdk.AccAddressFromBech32("cosmos14kve5tad89c6ewmae40cjjrn66gtgn6wkwsdhl")
		add2 := p.accountKeeper.GetAccount(p.ctx, add1)
		fmt.Println(add2.GetPubKey())
		fmt.Println(add2.GetAccountNumber())
		fmt.Println(add2.GetSequence())
		pbalances := p.bankkeeper.GetAccountsBalances(p.ctx)
		fmt.Println(pbalances)
		cosmosRpc := "http://127.0.0.1:26657"
		contact := NewContact(cosmosRpc)
		balances := contact.GetBalance("cosmos14kve5tad89c6ewmae40cjjrn66gtgn6wkwsdhl")
		fmt.Println("balances", balances)

	case ChainBSC:
	case ChainPI:
	case ChainICP:
	default:

	}

}

func (p *protocol) MonitorCrossTX(chain string) {

	switch chain {
	case ChainETH:
		chain, _ := p.chainKeeper.GetChain(p.ctx, chain)
		ethManager, _ := ethTools.NewEthManager(p.ctx, p.cdc, p.config, p.accountKeeper)
		ethManager.SyncLockDepositEvents(uint64(chain.Height))
	case ChainBSC:
	case ChainPI:
	case ChainICP:
	default:

	}
	// fmt.Println("cosmos1gyplaw49rstrxc2q9xhukml8vjkdvparz4r6y6")
	// add1, _ := sdk.AccAddressFromBech32("cosmos1gyplaw49rstrxc2q9xhukml8vjkdvparz4r6y6")
	// add2 := p.accountKeeper.GetAccount(p.ctx, add1)
	// fmt.Println(add2.GetPubKey())
	// fmt.Println(add2.GetAccountNumber())
	// fmt.Println(add2.GetSequence())
}
func (p *protocol) Start() {
	go p.MonitorChain(ChainETH)
	// go p.MonitorChain(ChainPI)
	// go p.MonitorCrossTX(ChainETH)
	// go p.MonitorCrossTX(ChainPI)
}

type ProtocolInstance interface {
	Start()
}

// func MakeProtocol(ctx sdk.Context, config *Config, accountKeeper authkeeper.AccountKeeper,
func MakeProtocol(ctx sdk.Context, cdc codec.Codec, config *config.ServiceConfig,
	accountKeeper authkeeper.AccountKeeper, bankkeeper bankkeeper.Keeper,
	blockKeeper blockmodulekeeper.Keeper, chainKeeper chainmodulekeeper.Keeper,
	nodeKeeper nodemodulekeeper.Keeper, trxKeeper trxmodulekeeper.Keeper,
) ProtocolInstance {
	p := &protocol{
		ctx:           ctx,
		cdc:           cdc,
		config:        config,
		accountKeeper: accountKeeper,
		bankkeeper:    bankkeeper,
		blockKeeper:   blockKeeper,
		chainKeeper:   chainKeeper,
		nodeKeeper:    nodeKeeper,
		trxKeeper:     trxKeeper,
	}

	return p
}
