package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

var BASE_URL = "https://api.beamlit.dev/v0"
var workspace string
var outputFormat string
var client *sdk.ClientWithResponses
var reg *Operations

var rootCmd = &cobra.Command{
	Use:   "beamlit",
	Short: "Beamlit CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if url := os.Getenv("BEAMLIT_API_URL"); url != "" {
			BASE_URL = url
		}

		reg = &Operations{
			BaseURL: BASE_URL,
		}

		provider := getAuthProvider(workspace)
		var err error
		client, err = sdk.NewClientWithResponses(
			BASE_URL,
			sdk.WithRequestEditorFn(provider.Intercept),
		)
		if err != nil {
			return err
		}

		ctx := context.Background()
		client.RegisterCliCommands(reg, ctx)
		return nil
	},
}

func Execute() error {
	reg = &Operations{
		BaseURL: BASE_URL,
	}

	rootCmd.AddCommand(reg.SetWorkspaceCmd())
	rootCmd.AddCommand(reg.GetWorkspaceCmd())
	rootCmd.AddCommand(reg.LoginCmd())
	rootCmd.AddCommand(reg.LogoutCmd())
	rootCmd.AddCommand(reg.GetCmd())
	rootCmd.AddCommand(reg.ApplyCmd())
	rootCmd.AddCommand(reg.DeleteCmd())

	rootCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "", "Specify the workspace name")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "Output format. One of: yaml")
	return rootCmd.Execute()
}
