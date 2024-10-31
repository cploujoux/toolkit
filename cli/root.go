package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

var BASE_URL = "https://api.beamlit.dev/v0"
var RUN_URL = "https://run.beamlit.dev"
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
		if runUrl := os.Getenv("BEAMLIT_RUN_URL"); runUrl != "" {
			RUN_URL = runUrl
		}

		reg = &Operations{
			BaseURL: BASE_URL,
			RunURL:  RUN_URL,
		}
		credentials := sdk.LoadCredentials(workspace)
		var err error
		c, err := sdk.NewClientWithCredentials(
			sdk.RunClientWithCredentials{
				ApiURL:      BASE_URL,
				RunURL:      RUN_URL,
				Credentials: credentials,
				Workspace:   workspace,
			},
		)
		if err != nil {
			return err
		}
		client = c

		ctx := context.Background()
		c.RegisterCliCommands(reg, ctx)
		return nil
	},
}

func Execute() error {
	reg = &Operations{
		BaseURL: BASE_URL,
	}

	rootCmd.AddCommand(reg.SetWorkspaceCmd())
	rootCmd.AddCommand(reg.GetWorkspaceCmd())
	rootCmd.AddCommand(reg.ListWorkspacesCmd())
	rootCmd.AddCommand(reg.LoginCmd())
	rootCmd.AddCommand(reg.LogoutCmd())
	rootCmd.AddCommand(reg.GetCmd())
	rootCmd.AddCommand(reg.ApplyCmd())
	rootCmd.AddCommand(reg.DeleteCmd())
	rootCmd.AddCommand(reg.RunCmd())
	rootCmd.AddCommand(reg.DocCmd())

	rootCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "", "Specify the workspace name")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "Output format. One of: yaml")
	if workspace == "" {
		workspace = sdk.CurrentWorkspace()
	}
	return rootCmd.Execute()
}
