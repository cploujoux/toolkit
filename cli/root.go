package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

var BASE_URL = "https://api.beamlit.dev/v0"
var workspace string

var provider *sdk.AuthProvider

var rootCmd = &cobra.Command{
	Use:   "beamlit",
	Short: "Beamlit CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		p := *provider
		p.SetWorkspace(workspace)
	},
}

func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "", "Specify the workspace name")
	ctx := context.Background()
	if url := os.Getenv("BEAMLIT_API_URL"); url != "" {
		BASE_URL = url
	}
	reg := &Operations{
		BaseURL: BASE_URL,
	}

	rootCmd.AddCommand(reg.SetWorkspaceCmd())
	rootCmd.AddCommand(reg.GetWorkspaceCmd())
	rootCmd.AddCommand(reg.LoginCmd())
	rootCmd.AddCommand(reg.LogoutCmd())
	rootCmd.AddCommand(reg.GetCmd())
	rootCmd.AddCommand(reg.ApplyCmd())
	rootCmd.AddCommand(reg.DeleteCmd())

	p := getAuthProvider()
	provider = &p

	client, err := sdk.NewClientWithResponses(
		BASE_URL,
		sdk.WithRequestEditorFn(p.Intercept),
	)
	if err != nil {
		return err
	}

	client.RegisterCliCommands(reg, ctx)

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
