package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	v1 "github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func txStart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [node-addr] [hours] [denom] [renewable]",
		Short: "Start a lease with a node for the specified duration and terms",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nodeAddr, err := base.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := strconv.ParseBool(args[3])
			if err != nil {
				return err
			}

			msg := v1.NewMsgStartRequest(
				ctx.FromAddress.Bytes(),
				nodeAddr,
				hours,
				args[2],
				renewable,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUpdateDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-details [id]",
		Short: "Update the renewable status of an existing lease",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := cmd.Flags().GetBool(flagRenewable)
			if err != nil {
				return err
			}

			msg := v1.NewMsgUpdateDetailsRequest(
				ctx.FromAddress.Bytes(),
				id,
				renewable,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Bool(flagRenewable, false, "specify if the lease is renewable")

	return cmd
}

func txRenew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew [id] [hours] [denom] [renewable]",
		Short: "Renew an existing lease for a specified duration and terms",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := strconv.ParseBool(args[3])
			if err != nil {
				return err
			}

			msg := v1.NewMsgRenewRequest(
				ctx.FromAddress.Bytes(),
				id,
				hours,
				args[2],
				renewable,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txEnd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end [id]",
		Short: "End an existing lease with the specified ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := v1.NewMsgEndRequest(
				ctx.FromAddress.Bytes(),
				id,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
