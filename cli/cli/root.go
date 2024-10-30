package cli

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tmp-moon/toolkit/sdk"
)

var BASE_URL = "https://api.beamlit.dev/v0"

func Execute() error {
	if url := os.Getenv("BEAMLIT_API_URL"); url != "" {
		BASE_URL = url
	}

	credentials := sdk.Credentials{}
	reg := &RegisterImpl{}

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

	ctx := context.Background()

	var provider sdk.AuthProvider
	if credentials.AccessToken != "" {
		provider = sdk.NewBearerTokenProvider(credentials.AccessToken, credentials.Workspace)
	}
	if credentials.APIKey != "" {
		provider = sdk.NewApiKeyProvider(credentials.APIKey, credentials.Workspace)
	}

	if provider != nil {
		client, err := sdk.NewClientWithResponses(
			BASE_URL,
			sdk.WithRequestEditorFn(provider.Intercept),
		)
		if err != nil {
			return err
		}

		client.RegisterCliCommands(reg, ctx)

		if err := execute(); err != nil {
			return err
		}
	}

	client, err := sdk.NewClientWithResponses(
		BASE_URL,
	)

	if err != nil {
		return err
	}

	client.RegisterCliCommands(reg, ctx)

	if err := execute(); err != nil {
		return err
	}

	return nil
}
