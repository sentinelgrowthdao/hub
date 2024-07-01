package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func querySubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription [subscription-id]",
		Short: "Query a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QuerySubscription(
				context.Background(),
				v2.NewQuerySubscriptionRequest(
					id,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func querySubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			accAddr, err := GetAccountAddress(cmd.Flags())
			if err != nil {
				return err
			}

			nodeAddr, err := GetNodeAddress(cmd.Flags())
			if err != nil {
				return err
			}

			planID, err := cmd.Flags().GetUint64(flagPlanID)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			if accAddr != nil {
				res, err := qc.QuerySubscriptionsForAccount(
					context.Background(),
					v2.NewQuerySubscriptionsForAccountRequest(
						accAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
			if nodeAddr != nil {
				res, err := qc.QuerySubscriptionsForNode(
					context.Background(),
					v2.NewQuerySubscriptionsForNodeRequest(
						nodeAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
			if planID != 0 {
				res, err := qc.QuerySubscriptionsForPlan(
					context.Background(),
					v2.NewQuerySubscriptionsForPlanRequest(
						planID,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QuerySubscriptions(
				context.Background(),
				v2.NewQuerySubscriptionsRequest(
					pagination,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "subscriptions")
	cmd.Flags().String(flagAccountAddress, "", "filter the subscriptions by an account address")
	cmd.Flags().String(flagNodeAddress, "", "filter the subscriptions by a node address")
	cmd.Flags().Uint64(flagPlanID, 0, "filter the subscriptions by a subscription plan")

	return cmd
}

func queryAllocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allocation [subscription-id] [account-addr]",
		Short: "Query an allocation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
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

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryAllocation(
				context.Background(),
				v2.NewQueryAllocationRequest(
					id,
					addr,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allocations [subscription-id]",
		Short: "Query allocations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryAllocations(
				context.Background(),
				v2.NewQueryAllocationsRequest(
					id,
					pagination,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "allocations")

	return cmd
}

func queryPayout() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payout [id]",
		Short: "Query a payout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryPayout(
				context.Background(),
				v2.NewQueryPayoutRequest(
					id,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryPayouts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payouts",
		Short: "Query payouts",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			accAddr, err := GetAccountAddress(cmd.Flags())
			if err != nil {
				return err
			}

			nodeAddr, err := GetNodeAddress(cmd.Flags())
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			if accAddr != nil {
				res, err := qc.QueryPayoutsForAccount(
					context.Background(),
					v2.NewQueryPayoutsForAccountRequest(
						accAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
			if nodeAddr != nil {
				res, err := qc.QueryPayoutsForNode(
					context.Background(),
					v2.NewQueryPayoutsForNodeRequest(
						nodeAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QueryPayouts(
				context.Background(),
				v2.NewQueryPayoutsRequest(
					pagination,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "payouts")
	cmd.Flags().String(flagAccountAddress, "", "filter the subscriptions by an account address")
	cmd.Flags().String(flagNodeAddress, "", "filter the subscriptions by a node address")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription-params",
		Short: "Query subscription module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryParams(
				context.Background(),
				v2.NewQueryParamsRequest(),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
