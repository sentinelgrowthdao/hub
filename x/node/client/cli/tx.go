// DO NOT COVER

package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func txRegister() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [remote-url] [gigabyte-prices] [hourly-prices]",
		Short: "Register a new node with the specified remote URL, gigabyte prices, and hourly prices",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gigabytePrices, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			hourlyPrices, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := v2.NewMsgRegisterRequest(
				ctx.FromAddress,
				gigabytePrices,
				hourlyPrices,
				args[0],
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
		Use:   "update-details",
		Short: "Update the pricing and remote URL details of an existing node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gigabytePrices, err := GetGigabytePrices(cmd.Flags())
			if err != nil {
				return err
			}

			hourlyPrices, err := GetHourlyPrices(cmd.Flags())
			if err != nil {
				return err
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := v2.NewMsgUpdateDetailsRequest(
				ctx.FromAddress.Bytes(),
				gigabytePrices,
				hourlyPrices,
				remoteURL,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagGigabytePrices, "", "prices per one gigabyte of bandwidth provision")
	cmd.Flags().String(flagHourlyPrices, "", "prices per one hour of bandwidth provision")
	cmd.Flags().String(flagRemoteURL, "", "remote URL address of the node")

	return cmd
}

func txUpdateStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-status [status]",
		Short: "Update the operational status of a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v2.NewMsgUpdateStatusRequest(
				ctx.FromAddress.Bytes(),
				v1base.StatusFromString(args[0]),
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

func txStartLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-lease [node-addr] [hours] [denom] [renewable]",
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

			msg := v3.NewMsgStartLeaseRequest(
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

func txUpdateLeaseDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-lease-details [id]",
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

			msg := v3.NewMsgUpdateLeaseDetailsRequest(
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

func txRenewLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew-lease [id] [hours] [denom] [renewable]",
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

			msg := v3.NewMsgRenewLeaseRequest(
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

func txEndLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end-lease [id]",
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

			msg := v3.NewMsgEndLeaseRequest(
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
