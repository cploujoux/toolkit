package beamlit

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var BASE_URL string

var rootCmd = &cobra.Command{
	Use:   "beamlit",
	Short: "Beamlit CLI",
}

func SetBaseURL(url string) {
	BASE_URL = url
}

func init() {
	if rootCmd.Use == "" {
		fmt.Println("Error: rootCmd not initialized")
		os.Exit(1)
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
}

func RegisterCliCommand(ctx context.Context, operationId string, fn interface{}) {
	operation := formatOperationId(operationId)
	if operation[0] == "list" {
		listRegister(ctx, operation, fn)
		return
	}
	if operation[0] == "get" {
		getRegister(ctx, operation, fn)
		return
	}
	if operation[0] == "delete" {
		deleteRegister(ctx, operation, fn)
		return
	}
	// fmt.Println("operation", operation)
	// if operation[0] == "u" {
	// 	applyRegister(ctx, operation, fn)
	// 	return
	// }
}

func Execute() error {
	return rootCmd.Execute()
}
