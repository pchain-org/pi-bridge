package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/pchain-org/pi-bridge/app"
	"github.com/spf13/viper"
	"github.com/tendermint/spm/cosmoscmd"
)

const (
	bscFlag      = "cross_bsc_enabled"
	dfinityFlag  = "cross_dfinity_enabled"
	ethereumFlag = "cross_ethereum_enabled"
	pchainFlag   = "cross_pchain_enabled"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.Flags().Bool(bscFlag, true, "support bsc cross-chain-transfer")
	rootCmd.Flags().Bool(ethereumFlag, true, "support ethereum cross-chain-transfer")
	rootCmd.Flags().Bool(dfinityFlag, true, "support dfinity cross-chain-transfer")
	rootCmd.Flags().Bool(pchainFlag, true, "support pchain cross-chain-transfer")
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		bscFlag := viper.GetBool(bscFlag)
		ethereumFlag := viper.GetBool(ethereumFlag)
		dfinityFlag := viper.GetBool(dfinityFlag)
		pchainFlag := viper.GetBool(pchainFlag)
		if bscFlag {
			fmt.Println("cross_bsc_enabled")
		}
		if ethereumFlag {
			fmt.Println("cross_ethereum_enabled")
		}
		if dfinityFlag {
			fmt.Println("cross_dfinity_enabled")
		}
		if pchainFlag {
			fmt.Println("cross_pchain_enabled")
		}

		os.Exit(1)
	}
}
