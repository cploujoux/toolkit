package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (r *Operations) DeployAgentAppCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "deploy-agent-app",
		Args:    cobra.ExactArgs(0),
		Aliases: []string{"da", "daa"},
		Short:   "Deploy a beamlit agent app",
		Long:    "Deploy a beamlit agent app, you must be in a beamlit agent app directory.",
		Example: `bl deploy-agent-app`,
		Run: func(cmd *cobra.Command, args []string) {

			agentName := "test"
			fmt.Printf(`Your beamlit agent app has been deployed. Check it out at %s/%s/%s`, r.AppURL, workspace, agentName)
		},
	}
	return cmd
}
