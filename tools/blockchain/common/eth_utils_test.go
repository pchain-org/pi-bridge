package common

import (
	"fmt"
	"testing"

	"github.com/pchain-org/pi-bridge/tools/config"
)

var servConfig *config.ServiceConfig

func init() {
	servConfig = config.NewServiceConfig("./../../../cs_config.json")
	if servConfig == nil {
		fmt.Print("test init error")
	}
}

func TestGetEthHeight(t *testing.T) {
	url := servConfig.ETHConfig.RestURL
	restClient := NewRestClient()
	heigth, err := GetEthHeight(url, restClient)
	if err != nil {
		t.Log(err)
	}
	t.Log(heigth)
}
