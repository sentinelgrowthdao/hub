package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

func queryProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider [provider-addr]",
		Short: "Query a provider",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := base.ProvAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var (
				qc = v2.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryProvider(
				context.Background(),
				v2.NewQueryProviderRequest(
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

func queryProviders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "providers",
		Short: "Query providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
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

			res, err := qc.QueryProviders(
				context.Background(),
				v2.NewQueryProvidersRequest(
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
	flags.AddPaginationFlagsToCmd(cmd, "providers")
	cmd.Flags().String(base.FlagStatus, "", "filter the providers by status (active|inactive)")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider-params",
		Short: "Query provider module parameters",
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
