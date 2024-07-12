package cli

import (
	"strconv"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func txCancelSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-subscription [id]",
		Short: "Cancel an active subscription",
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

			msg := v3.NewMsgCancelSubscriptionRequest(
				ctx.FromAddress,
				id,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txRenewSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew-subscription [id] [denom]",
		Short: "Renew an existing subscription",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := v3.NewMsgRenewSubscriptionRequest(
				ctx.FromAddress.Bytes(),
				id,
				args[1],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txShareSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "share-subscription [id] [acc-addr] [bytes]",
		Short: "Share a subscription with an account",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			bytes, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := v3.NewMsgShareSubscriptionRequest(
				ctx.FromAddress.Bytes(),
				id,
				addr,
				sdkmath.NewInt(bytes),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txStartSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-subscription [id] [denom] [renewable]",
		Short: "Start a subscription for a plan",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}

			msg := v3.NewMsgStartSubscriptionRequest(
				ctx.FromAddress.Bytes(),
				id,
				args[1],
				renewable,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUpdateSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-subscription [id] [renewable]",
		Short: "Update the details of an existing subscription",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdateSubscriptionRequest(
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

	return cmd
}

func txStartSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-session [id] [node-addr]",
		Short: "Start a session for a subscription and node",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			nodeAddr, err := base.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := v3.NewMsgStartSessionRequest(
				ctx.FromAddress.Bytes(),
				id,
				nodeAddr,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
