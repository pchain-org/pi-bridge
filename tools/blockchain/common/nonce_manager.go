package common

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

const clear_nonce_interval = 10 * time.Minute

type NonceManager struct {
	addressNonce  map[common.Address]uint64
	returnedNonce map[common.Address]SortedNonceArr
	ethClient     *ethclient.Client
	lock          sync.Mutex
}

func NewNonceManager(ethClient *ethclient.Client) *NonceManager {
	nonceManager := &NonceManager{
		addressNonce:  make(map[common.Address]uint64),
		ethClient:     ethClient,
		returnedNonce: make(map[common.Address]SortedNonceArr),
	}
	//go nonceManager.clearNonce()
	return nonceManager
}

// return account nonce, and than nonce++
func (nm *NonceManager) GetAddressNonce(address common.Address) uint64 {
	nm.lock.Lock()
	defer nm.lock.Unlock()

	if nm.returnedNonce[address].Len() > 0 {
		nonce := nm.returnedNonce[address][0]
		nm.returnedNonce[address] = nm.returnedNonce[address][1:]
		return nonce
	}

	// return a new point
	nonce, ok := nm.addressNonce[address]
	if !ok {
		// get nonce from eth network
		uintNonce, err := nm.ethClient.PendingNonceAt(context.Background(), address)
		if err != nil {
			log.NewNopLogger().Error("GetAddressNonce: cannot get account %s nonce, err: %s, set it to nil!",
				address, err)
		}
		nm.addressNonce[address] = uintNonce
		nonce = uintNonce
	}
	// increase record
	nm.addressNonce[address]++
	return nonce
}

func (nm *NonceManager) ReturnNonce(addr common.Address, nonce uint64) {
	nm.lock.Lock()
	defer nm.lock.Unlock()

	arr, ok := nm.returnedNonce[addr]
	if !ok {
		arr = make([]uint64, 0)
	}
	arr = append(arr, nonce)
	sort.Sort(arr)
	nm.returnedNonce[addr] = arr
}

func (nm *NonceManager) DecreaseAddressNonce(address common.Address) {
	nm.lock.Lock()
	defer nm.lock.Unlock()

	nonce, ok := nm.addressNonce[address]
	if ok && nonce > 0 {
		nm.addressNonce[address]--
	}
}

// clear nonce per
func (nm *NonceManager) clearNonce() {
	for {
		<-time.After(clear_nonce_interval)
		nm.lock.Lock()
		for addr, _ := range nm.addressNonce {
			delete(nm.addressNonce, addr)
		}
		nm.lock.Unlock()
		//log.Infof("clearNonce: clear all cache nonce")
	}
}

type SortedNonceArr []uint64

func (arr SortedNonceArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr SortedNonceArr) Len() int { return len(arr) }

func (arr SortedNonceArr) Swap(i, j int) { arr[i], arr[j] = arr[j], arr[i] }
