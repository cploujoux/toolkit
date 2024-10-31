package operations

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) GetContextCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get-context",
		Short: "Get the current workspace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(sdk.CurrentContext())
		},
	}
}

func (r *Operations) SetContextCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-context [workspace]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Set the current workspace",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Error: Workspace is required")
				os.Exit(1)
			} else {
				sdk.SetCurrentContext(args[0])
			}
		},
	}
}
