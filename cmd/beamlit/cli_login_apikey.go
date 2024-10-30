package beamlit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func apiKeyLogin(workspace string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	credentialsDir := filepath.Join(homeDir, ".beamlit")
	credentialsFile := filepath.Join(credentialsDir, "credentials.json")

	// Check if credentials file exists
	if _, err := os.Stat(credentialsFile); err == nil {
		fmt.Println("You are already logged in. Enter new API key to overwrite")
	} else {
		if err := os.MkdirAll(credentialsDir, 0700); err != nil {
			fmt.Printf("Error creating credentials directory: %v\n", err)
			return
		}
		fmt.Println("Enter your API key")
	}

	var apiKey string
	for {
		fmt.Scanln(&apiKey)

		if apiKey != "" {
			break
		}
		fmt.Println("API key cannot be empty. Please enter your API key")
	}

	// Create credentials struct and marshal to JSON
	credentials := Credentials{
		APIKey:    apiKey,
		Workspace: workspace,
	}

	jsonData, err := json.MarshalIndent(credentials, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling credentials: %v\n", err)
		return
	}

	// Write JSON to credentials file
	if err := os.WriteFile(credentialsFile, jsonData, 0600); err != nil {
		fmt.Printf("Error writing credentials file: %v\n", err)
		return
	}

	fmt.Println("Successfully stored API key")
}
