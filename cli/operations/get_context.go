package operations

import (
	"fmt"

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
