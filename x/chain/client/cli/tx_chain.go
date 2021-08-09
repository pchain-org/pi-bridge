package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pchain-org/pi-bridge/x/chain/types"
)

func CmdCreateChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [index] [sourceChainID] [height] [proof] [proxyAddress] [extra] [headerOrCrossChainMsg]",
		Short: "Create a new Chain",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]
			argsSourceChainID, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argsHeight, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argsProof, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsProxyAddress, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsExtra, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsHeaderOrCrossChainMsg, err := cast.ToStringE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateChain(clientCtx.GetFromAddress().String(), index, argsSourceChainID, argsHeight, argsProof, argsProxyAddress, argsExtra, argsHeaderOrCrossChainMsg)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-chain [index] [sourceChainID] [height] [proof] [proxyAddress] [extra] [headerOrCrossChainMsg]",
		Short: "Update a Chain",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			argsSourceChainID, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argsHeight, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argsProof, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsProxyAddress, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsExtra, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsHeaderOrCrossChainMsg, err := cast.ToStringE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateChain(clientCtx.GetFromAddress().String(), index, argsSourceChainID, argsHeight, argsProof, argsProxyAddress, argsExtra, argsHeaderOrCrossChainMsg)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-chain [index]",
		Short: "Delete a Chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteChain(clientCtx.GetFromAddress().String(), index)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
