package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) ListWorkspacesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get-workspaces",
		Short: "List all workspaces",
		Run: func(cmd *cobra.Command, args []string) {
			workspaces := sdk.ListWorkspaces()
			currentWorkspace := sdk.CurrentWorkspace()
			for _, workspace := range workspaces {
				if workspace == currentWorkspace {
					fmt.Printf("* %s\n", workspace)
				} else {
					fmt.Printf("  %s\n", workspace)
				}
			}
		},
	}
}

func (r *Operations) GetWorkspaceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get-workspace",
		Short: "Get the current workspace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(sdk.CurrentWorkspace())
		},
	}
}

func (r *Operations) SetWorkspaceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-workspace [workspace]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Set the current workspace",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Error: Workspace is required")
				os.Exit(1)
			} else {
				sdk.SetCurrentWorkspace(args[0])
			}
		},
	}
}
