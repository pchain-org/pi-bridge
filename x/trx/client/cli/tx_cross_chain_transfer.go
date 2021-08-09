package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

var _ = strconv.Itoa(0)

func CmdCrossChainTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cross-chain-transfer [asset] [amount] [fromChain] [fromAddress] [toChain] [toAddress]",
		Short: "Broadcast message CrossChainTransfer",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAsset := string(args[0])
			argsAmount, _ := strconv.ParseInt(args[1], 10, 64)
			argsFromChain := string(args[2])
			argsFromAddress := string(args[3])
			argsToChain := string(args[4])
			argsToAddress := string(args[5])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCrossChainTransfer(clientCtx.GetFromAddress().String(), string(argsAsset), int32(argsAmount), string(argsFromChain), string(argsFromAddress), string(argsToChain), string(argsToAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
