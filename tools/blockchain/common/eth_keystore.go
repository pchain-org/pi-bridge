package common

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pchain-org/pi-bridge/tools/config"
	"github.com/tendermint/tendermint/libs/log"
)

type EthKeyStore struct {
	ks      *keystore.KeyStore
	chainId *big.Int
}

func NewEthKeyStore(sigConfig *config.ETHConfig, chainId *big.Int) *EthKeyStore {
	service := &EthKeyStore{}
	capitalKeyStore := keystore.NewKeyStore(sigConfig.KeyStorePath, keystore.StandardScryptN,
		keystore.StandardScryptP)
	accArr := capitalKeyStore.Accounts()
	if len(accArr) == 0 {
		log.NewNopLogger().Error("relayer has no account")
		panic(fmt.Errorf("relayer has no account"))
	}
	str := ""
	for i, v := range accArr {
		str += fmt.Sprintf("(no.%d acc: %s), ", i+1, v.Address.String())
	}
	log.NewNopLogger().Info("relayer are using accounts: [ %s ]", str)
	service.ks = capitalKeyStore
	service.chainId = chainId
	return service
}

func (ks *EthKeyStore) UnlockKeys(sigConfig *config.ETHConfig) error {
	for _, v := range ks.GetAccounts() {
		err := ks.ks.Unlock(v, sigConfig.KeyStorePwdSet[strings.ToLower(v.Address.String())])
		if err != nil {
			return fmt.Errorf("failed to unlock eth acc %s: %v", v.Address.String(), err)
		}
	}
	return nil
}

func (ks *EthKeyStore) SignTransaction(tx *types.Transaction, acc accounts.Account) (*types.Transaction, error) {
	tx, err := ks.ks.SignTx(acc, tx, ks.chainId)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (ks *EthKeyStore) GetAccounts() []accounts.Account {
	return ks.ks.Accounts()
}

func (ks *EthKeyStore) TestPwd(acc accounts.Account, pwd string) error {
	if err := ks.ks.Unlock(acc, pwd); err != nil {
		return err
	}
	_ = ks.ks.Lock(acc.Address)
	return nil
}

func (ks *EthKeyStore) GetChainId() uint64 {
	return ks.chainId.Uint64()
}
