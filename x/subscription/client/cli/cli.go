package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		querySubscription(),
		querySubscriptions(),
		queryAllocation(),
		queryAllocations(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Subscription module sub-commands",
	}

	cmd.AddCommand(
		txCancel(),
		txAllocate(),
		txStart(),
		txUpdate(),
		txRenew(),
		txStartSession(),
	)

	return []*cobra.Command{cmd}
}
