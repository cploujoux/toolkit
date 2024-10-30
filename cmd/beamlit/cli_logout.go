package beamlit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Beamlit",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			return
		}

		credentialsFile := filepath.Join(homeDir, ".beamlit", "credentials.json")
		if err := os.Remove(credentialsFile); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Already logged out")
				return
			}
			fmt.Printf("Error removing credentials file: %v\n", err)
			return
		}

		fmt.Println("Successfully logged out")
	},
}
