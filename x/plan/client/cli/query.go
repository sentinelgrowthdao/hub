// DO NOT COVER

package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

func queryPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan [id]",
		Short: "Query a plan",
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

			res, err := qc.QueryPlan(
				context.Background(),
				v2.NewQueryPlanRequest(
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

func queryPlans() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plans",
		Short: "Query plans",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			provAddr, err := GetProvider(cmd.Flags())
			if err != nil {
				return err
			}

			status, err := GetStatus(cmd.Flags())
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

			if provAddr != nil {
				res, err := qc.QueryPlansForProvider(
					context.Background(),
					v2.NewQueryPlansForProviderRequest(
						provAddr,
						status,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QueryPlans(
				context.Background(),
				v2.NewQueryPlansRequest(
					status,
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
	flags.AddPaginationFlagsToCmd(cmd, "plans")
	cmd.Flags().String(flagProvider, "", "filter the plans by provider address")
	cmd.Flags().String(flagStatus, "", "filter the plans by status (active|inactive)")

	return cmd
}
