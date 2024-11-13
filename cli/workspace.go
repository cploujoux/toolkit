package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) ListOrSetWorkspacesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "workspaces [workspace]",
		Aliases: []string{"ws", "workspace"},
		Short:   "List all workspaces with the current workspace highlighted, set optionally a new current workspace",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				sdk.SetCurrentWorkspace(args[0])
			}

			workspaces := sdk.ListWorkspaces()
			currentWorkspace := sdk.CurrentWorkspace()

			// Afficher les en-têtes du tableau
			fmt.Printf("%-20s %s\n", "NAME", "CURRENT")

			// Afficher chaque workspace
			for _, workspace := range workspaces {
				current := " "
				if workspace == currentWorkspace {
					current = "*"
				}
				fmt.Printf("%-20s %s\n", workspace, current)
			}
		},
	}
}
