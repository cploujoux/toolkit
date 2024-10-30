package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmp-moon/toolkit/cli/operations"
)

var rootCmd = &cobra.Command{
	Use:   "beamlit",
	Short: "Beamlit CLI",
}

func init() {
	if rootCmd.Use == "" {
		fmt.Println("Error: rootCmd not initialized")
		os.Exit(1)
	}

	operations.SetBaseURL(BASE_URL)

	rootCmd.AddCommand(operations.ListCmd)
	rootCmd.AddCommand(operations.GetCmd)
	rootCmd.AddCommand(operations.DeleteCmd)
	rootCmd.AddCommand(operations.ApplyCmd)
	rootCmd.AddCommand(operations.LoginCmd)
	rootCmd.AddCommand(operations.LogoutCmd)
}

type RegisterImpl struct {
}

func (r *RegisterImpl) CliCommand(ctx context.Context, operationId string, fn interface{}) {
	operation := formatOperationId(operationId)
	if operation[0] == "list" {
		operations.ListRegister(ctx, operation, fn)
		return
	}
	if operation[0] == "get" {
		operations.GetRegister(ctx, operation, fn)
		return
	}
	if operation[0] == "delete" {
		operations.DeleteRegister(ctx, operation, fn)
		return
	}
	// fmt.Println("operation", operation)
	// if operation[0] == "u" {
	// 	applyRegister(ctx, operation, fn)
	// 	return
	// }

}

func execute() error {
	return rootCmd.Execute()
}
