// DO NOT COVER

package cli

import (
	"context"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v12/x/swap/types"
	"github.com/sentinel-official/hub/v12/x/swap/types/v1"
)

func querySwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap",
		Short: "Query a swap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			txHash, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			var (
				qc = v1.NewQueryServiceClient(ctx)
			)

			res, err := qc.QuerySwap(
				context.Background(),
				v1.NewQuerySwapRequest(
					types.BytesToHash(txHash),
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

func querySwaps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaps",
		Short: "Query swaps",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
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

			res, err := qc.QuerySwaps(
				context.Background(),
				v1.NewQuerySwapsRequest(
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
	flags.AddPaginationFlagsToCmd(cmd, "swaps")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query swap module parameters",
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
				v1.NewQueryParamsRequest(),
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
