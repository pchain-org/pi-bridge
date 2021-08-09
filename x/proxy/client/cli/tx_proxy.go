package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pchain-org/pi-bridge/x/proxy/types"
)

func CmdCreateProxy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-proxy [index] [addressList] [address]",
		Short: "Create a new Proxy",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]
			argsAddressList, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsAddress, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProxy(clientCtx.GetFromAddress().String(), index, argsAddressList, argsAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateProxy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-proxy [index] [addressList] [address]",
		Short: "Update a Proxy",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			argsAddressList, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsAddress, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProxy(clientCtx.GetFromAddress().String(), index, argsAddressList, argsAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteProxy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-proxy [index]",
		Short: "Delete a Proxy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProxy(clientCtx.GetFromAddress().String(), index)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
