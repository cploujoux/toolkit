package cli

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tmp-moon/toolkit/sdk"
)

func loadCredentials() sdk.Credentials {
	credentials := sdk.Credentials{}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		credentialsPath := filepath.Join(homeDir, ".beamlit", "credentials.json")
		if data, err := os.ReadFile(credentialsPath); err == nil {
			if err := json.Unmarshal(data, &credentials); err != nil {
				// Invalid JSON, use empty credentials
				credentials = sdk.Credentials{}
			}
		}
	}

	return credentials
}

func getAuthProvider() sdk.AuthProvider {
	credentials := loadCredentials()

	var provider sdk.AuthProvider
	if credentials.AccessToken != "" {
		provider = sdk.NewBearerTokenProvider(credentials.AccessToken, credentials.Workspace)
	} else if credentials.APIKey != "" {
		provider = sdk.NewApiKeyProvider(credentials.APIKey, credentials.Workspace)
	} else {
		provider = sdk.NewPublicProvider()
	}

	return provider
}
