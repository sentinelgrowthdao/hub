// DO NOT COVER

package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func queryNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node [node-addr]",
		Short: "Query a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := base.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryNode(
				context.Background(),
				v2.NewQueryNodeRequest(
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

func queryNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Query nodes",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := cmd.Flags().GetUint64(flagPlanID)
			if err != nil {
				return err
			}

			status, err := base.StatusFromFlags(cmd.Flags())
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

			switch {
			case id != 0:
				res, err := qc.QueryNodesForPlan(
					context.Background(),
					v2.NewQueryNodesForPlanRequest(
						id,
						status,
						pagination,
					),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			default:
				res, err := qc.QueryNodes(
					context.Background(),
					v2.NewQueryNodesRequest(
						status,
						pagination,
					),
				)
				if err != nil {
					return err
				}
				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nodes")
	cmd.Flags().String(base.FlagStatus, "", "filter the nodes by status (active|inactive)")
	cmd.Flags().Uint64(flagPlanID, 0, "filter the nodes by subscription plan ID")

	return cmd
}

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
				qc = v3.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryLease(
				context.Background(),
				&v3.QueryLeaseRequest{
					ID: id,
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
				qc = v3.NewQueryServiceClient(ctx)
			)

			switch {
			case nodeAddr != "":
				res, err := qc.QueryLeasesForNode(
					context.Background(),
					&v3.QueryLeasesForNodeRequest{
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
					&v3.QueryLeasesForProviderRequest{
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
					&v3.QueryLeasesRequest{
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
		Use:   "node-params",
		Short: "Query node module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				qc = v3.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryParams(
				context.Background(),
				&v3.QueryParamsRequest{},
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
