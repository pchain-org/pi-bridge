package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pchain-org/pi-bridge/x/node/types"
)

func CmdCreateNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-node [index] [peerPubkey] [address]",
		Short: "Create a new Node",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]
			argsPeerPubkey, err := cast.ToStringE(args[1])
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

			msg := types.NewMsgCreateNode(clientCtx.GetFromAddress().String(), index, argsPeerPubkey, argsAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-node [index] [peerPubkey] [address]",
		Short: "Update a Node",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			argsPeerPubkey, err := cast.ToStringE(args[1])
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

			msg := types.NewMsgUpdateNode(clientCtx.GetFromAddress().String(), index, argsPeerPubkey, argsAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-node [index]",
		Short: "Delete a Node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteNode(clientCtx.GetFromAddress().String(), index)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
