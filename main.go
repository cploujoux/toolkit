package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/tmp-moon/toolkit/cmd/beamlit"
)

var BASE_URL = "https://api.beamlit.dev/v0"

func init() {
	if url := os.Getenv("BEAMLIT_API_URL"); url != "" {
		BASE_URL = url
	}
}

func main() {
	beamlit.SetBaseURL(BASE_URL)
	credentials := beamlit.Credentials{}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		credentialsPath := filepath.Join(homeDir, ".beamlit", "credentials.json")
		if data, err := os.ReadFile(credentialsPath); err == nil {
			if err := json.Unmarshal(data, &credentials); err != nil {
				// Invalid JSON, use empty credentials
				credentials = beamlit.Credentials{}
			}
		}
	}

	ctx := context.Background()

	var provider beamlit.AuthProvider
	if credentials.AccessToken != "" {
		provider = beamlit.NewBearerTokenProvider(credentials.AccessToken, credentials.Workspace)
	}
	if credentials.APIKey != "" {
		provider = beamlit.NewApiKeyProvider(credentials.APIKey, credentials.Workspace)
	}

	if provider != nil {
		client, err := beamlit.NewClientWithResponses(
			BASE_URL,
			beamlit.WithRequestEditorFn(provider.Intercept),
		)
		if err != nil {
			slog.Error("Error creating client", "error", err)
			os.Exit(1)
		}

		client.RegisterCliCommands(ctx)

		if err := beamlit.Execute(); err != nil {
			slog.Error("Error executing command", "error", err)
			os.Exit(1)
		}
		return
	}

	client, err := beamlit.NewClientWithResponses(
		BASE_URL,
	)

	if err != nil {
		slog.Error("Error creating client", "error", err)
		os.Exit(1)
	}

	client.RegisterCliCommands(ctx)

	if err := beamlit.Execute(); err != nil {
		slog.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}
