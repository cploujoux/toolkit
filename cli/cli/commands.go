package cli

import (
	"fmt"
	"os"
)

func init() {
	if rootCmd.Use == "" {
		fmt.Println("Error: rootCmd not initialized")
		os.Exit(1)
	}
}

type RegisterImpl struct {
}
