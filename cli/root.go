package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

var BASE_URL = "https://api.beamlit.dev/v0"
var workspaceFlag string

var rootCmd = &cobra.Command{
	Use:   "beamlit",
	Short: "Beamlit CLI",
}

func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&workspaceFlag, "workspace", "w", "", "Specify the workspace name")

	ctx := context.Background()
	if url := os.Getenv("BEAMLIT_API_URL"); url != "" {
		BASE_URL = url
	}
	reg := &Operations{
		BaseURL: BASE_URL,
	}

	for _, cmd := range reg.MainCommand() {
		rootCmd.AddCommand(cmd)
	}

	provider := getAuthProvider(workspaceFlag)

	client, err := sdk.NewClientWithResponses(
		BASE_URL,
		sdk.WithRequestEditorFn(provider.Intercept),
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
