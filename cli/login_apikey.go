package cli

import (
	"fmt"

	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) ApiKeyLogin(workspace string) {
	var apiKey string
	for {
		fmt.Scanln(&apiKey)

		if apiKey != "" {
			break
		}
		fmt.Println("API key cannot be empty. Please enter your API key")
	}

	// Create credentials struct and marshal to JSON
	credentials := sdk.Credentials{
		APIKey: apiKey,
	}

	sdk.SaveCredentials(workspace, credentials)
	fmt.Println("Successfully stored API key")
}
