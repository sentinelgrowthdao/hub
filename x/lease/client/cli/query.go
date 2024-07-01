package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func queryLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease [id]",
		Short: "Query a lease",
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
				qc = v1.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryLease(
				context.Background(),
				&v1.QueryLeaseRequest{
					Id: id,
				},
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

func queryLeases() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leases",
		Short: "Query leases with optional filters for node or provider",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nodeAddr, err := cmd.Flags().GetString(flagNodeAddr)
			if err != nil {
				return err
			}

			provAddr, err := cmd.Flags().GetString(flagProvAddr)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var (
				qc = v1.NewQueryServiceClient(ctx)
			)

			switch {
			case nodeAddr != "":
				res, err := qc.QueryLeasesForNode(
					context.Background(),
					&v1.QueryLeasesForNodeRequest{
						Address:    nodeAddr,
						Pagination: pagination,
					},
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			case provAddr != "":
				res, err := qc.QueryLeasesForProvider(
					context.Background(),
					&v1.QueryLeasesForProviderRequest{
						Address:    provAddr,
						Pagination: pagination,
					},
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			default:
				res, err := qc.QueryLeases(
					context.Background(),
					&v1.QueryLeasesRequest{
						Pagination: pagination,
					},
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "leases")
	cmd.Flags().String(flagNodeAddr, "", "filter the leases by node address")
	cmd.Flags().String(flagProvAddr, "", "filter the leases by provider address")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease-params",
		Short: "Query lease module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				qc = v1.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryParams(
				context.Background(),
				&v1.QueryParamsRequest{},
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
