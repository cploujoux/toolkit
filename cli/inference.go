package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) InferenceCmd() *cobra.Command {
	var data string
	cmd := &cobra.Command{
		Use:     "inference [model] [environment]",
		Args:    cobra.MaximumNArgs(2),
		Short:   "Inference",
		Aliases: []string{"infer"},
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
				sdk.Inference(context.Background(), workspace, args[1], args[0], data)
			} else {
				sdk.Inference(context.Background(), workspace, "production", args[0], data)
			}
		},
	}

	cmd.Flags().StringVar(&data, "data", "", "JSON body data for the inference request")
	return cmd
}
