package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pchain-org/pi-bridge/x/block/types"
)

func CmdCreateBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-block [index] [chainID] [address] [headers]",
		Short: "Create a new Block",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]
			argsChainID, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argsAddress, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsHeaders, err := cast.ToStringSliceE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBlock(clientCtx.GetFromAddress().String(), index, argsChainID, argsAddress, argsHeaders)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-block [index] [chainID] [address] [headers]",
		Short: "Update a Block",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			argsChainID, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argsAddress, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsHeaders, err := cast.ToStringSliceE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBlock(clientCtx.GetFromAddress().String(), index, argsChainID, argsAddress, argsHeaders)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-block [index]",
		Short: "Delete a Block",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteBlock(clientCtx.GetFromAddress().String(), index)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
