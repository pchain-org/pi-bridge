package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pchain-org/pi-bridge/x/trx/types"
)

func CmdCreateTrx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-trx [index] [chainID] [address] [crossChainMsgs]",
		Short: "Create a new Trx",
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
			argsCrossChainMsgs, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			proof, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTrx(clientCtx.GetFromAddress().String(), index, argsChainID, argsAddress, argsCrossChainMsgs, proof)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateTrx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-trx [index] [chainID] [address] [crossChainMsgs]",
		Short: "Update a Trx",
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
			argsCrossChainMsgs, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTrx(clientCtx.GetFromAddress().String(), index, argsChainID, argsAddress, argsCrossChainMsgs)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteTrx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-trx [index]",
		Short: "Delete a Trx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteTrx(clientCtx.GetFromAddress().String(), index)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
