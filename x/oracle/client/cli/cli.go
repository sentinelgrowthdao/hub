package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle",
		Short: "Querying commands for the Oracle module",
	}

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle",
		Short: "Oracle transactions subcommands",
	}

	return cmd
}
