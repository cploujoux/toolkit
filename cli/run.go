package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Operations) RunCmd() *cobra.Command {
	var data string
	cmd := &cobra.Command{
		Use:   "run [model] [environment]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Run inference",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Error: Model is required")
				os.Exit(1)
			}

			if data == "" {
				fmt.Println("Error: --data parameter is required")
				os.Exit(1)
			}

			if len(args) == 2 {
				client.Run(context.Background(), workspace, args[1], args[0], data)
			} else {
				client.Run(context.Background(), workspace, "production", args[0], data)
			}
		},
	}

	cmd.Flags().StringVar(&data, "data", "", "JSON body data for the inference request")
	return cmd
}
