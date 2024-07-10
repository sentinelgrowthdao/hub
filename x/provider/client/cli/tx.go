package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

func txRegister() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [name]",
		Short: "Register a new provider with a name and optional details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			identity, err := cmd.Flags().GetString(flagIdentity)
			if err != nil {
				return err
			}

			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				return err
			}

			msg := v2.NewMsgRegisterRequest(
				ctx.FromAddress.Bytes(),
				args[0],
				identity,
				website,
				description,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagIdentity, "", "unique identity of the provider")
	cmd.Flags().String(flagWebsite, "", "official website URL of the provider")
	cmd.Flags().String(flagDescription, "", "brief description of the provider's services or offerings")

	return cmd
}

func txUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the details of an existing provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString(flagName)
			if err != nil {
				return err
			}

			identity, err := cmd.Flags().GetString(flagIdentity)
			if err != nil {
				return err
			}

			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				return err
			}

			status, err := base.StatusFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			msg := v2.NewMsgUpdateRequest(
				ctx.FromAddress.Bytes(),
				name,
				identity,
				website,
				description,
				status,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagName, "", "name of the provider")
	cmd.Flags().String(flagIdentity, "", "unique identity of the provider")
	cmd.Flags().String(flagWebsite, "", "official website URL of the provider")
	cmd.Flags().String(flagDescription, "", "breif description of the provider's services or offerings")
	cmd.Flags().String(base.FlagStatus, "", "operational status of the provider (active|inactive)")

	return cmd
}
