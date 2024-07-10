package cli

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

func txCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [duration] [gigabytes] [prices]",
		Short: "Create a new subscription plan with duration, gigabytes and pricing details",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			duration, err := time.ParseDuration(args[0])
			if err != nil {
				return err
			}

			gigabytes, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			prices, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := v2.NewMsgCreateRequest(
				ctx.FromAddress.Bytes(),
				duration,
				gigabytes,
				prices,
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

func txUpdateStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-status [id] [status]",
		Short: "Update the status of an existing subscription plan",
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

			msg := v2.NewMsgUpdateStatusRequest(
				ctx.FromAddress.Bytes(),
				id,
				v1base.StatusFromString(args[1]),
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

func txLinkNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-node [id] [node-addr]",
		Short: "Link a node to a subscription plan",
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

			addr, err := base.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := v2.NewMsgLinkNodeRequest(
				ctx.FromAddress.Bytes(),
				id,
				addr,
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

func txUnlinkNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink-node [id] [node-addr]",
		Short: "Unlink a node from a subscription plan",
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

			addr, err := base.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := v2.NewMsgUnlinkNodeRequest(
				ctx.FromAddress.Bytes(),
				id,
				addr,
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
