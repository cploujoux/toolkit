package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) LogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout [workspace]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Logout from Beamlit",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Error: Workspace is required")
				os.Exit(1)
			} else {
				sdk.ClearCredentials(args[0])
			}
		},
	}
}
