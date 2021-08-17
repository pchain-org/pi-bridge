package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/tendermint/tendermint/libs/log"
)

const (
	ETH_MONITOR_INTERVAL    = 15 * time.Second
	BEIDGE_MONITOR_INTERVAL = 1 * time.Second

	ETH_USEFUL_BLOCK_NUM     = 12
	ETH_PROOF_USERFUL_BLOCK  = 12
	PI_USEFUL_BLOCK_NUM      = 1
	DEFAULT_CONFIG_FILE_NAME = "./cs_config.json"
	Version                  = "1.0"
)

type ServiceConfig struct {
	BridgeConfig *BridgeConfig
	ETHConfig    *ETHConfig
	BoltDbPath   string
	RoutineNum   int64
	// TargetContracts []map[string]map[string][]uint64
}

type BridgeConfig struct {
	BridgeRpcAddr        string
	BridgeWallet         string
	BridgeWalletPwd      string
	BridgeStartHeight    int64
	BridgeListenInterval int
	BridgeChainId        string
	BridgeGasPrice       string
	BridgeGas            uint64
}

type ETHConfig struct {
	SideChainId         uint64
	RestURL             string
	ECCMContractAddress string
	ECCDContractAddress string
	KeyStorePath        string
	KeyStorePwdSet      map[string]string
	BlockConfig         uint64
	HeadersPerBatch     int
	MonitorInterval     uint64
	TargetContracts     []map[string]map[string][]uint64
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.NewNopLogger().Error("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func NewServiceConfig(configFilePath string) *ServiceConfig {
	fileContent, err := ReadFile(configFilePath)
	if err != nil {
		log.NewNopLogger().Error("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	servConfig := &ServiceConfig{}
	err = json.Unmarshal(fileContent, servConfig)
	if err != nil {
		log.NewNopLogger().Error("NewServiceConfig: failed, err: %s", err)
		return nil
	}

	for k, v := range servConfig.ETHConfig.KeyStorePwdSet {
		delete(servConfig.ETHConfig.KeyStorePwdSet, k)
		servConfig.ETHConfig.KeyStorePwdSet[strings.ToLower(k)] = v
	}

	return servConfig
}
