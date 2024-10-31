package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/tmp-moon/toolkit/sdk"
)

var BASE_URL = "https://api.sdk.dev/v0"

func init() {
	if url := os.Getenv("BEAMLIT_API_URL"); url != "" {
		BASE_URL = url
	}
}
func main() {
	ctx := context.Background()

	provider := sdk.NewApiKeyProvider(sdk.Credentials{APIKey: "BL_96UX6OFWA6JG9IA1NDE7CCMED22U8Z2F"}, "chris")

	client, err := sdk.NewClientWithResponses(
		BASE_URL,
		sdk.WithRequestEditorFn(provider.Intercept),
	)
	if err != nil {
		slog.Error("Error creating client", "error", err)
		os.Exit(1)
	}

	env, err := client.PutEnvironmentWithResponse(ctx, "test", sdk.PutEnvironmentJSONRequestBody{
		Name:        sdk.BlString("test"),
		DisplayName: sdk.BlString("Test"),
	})
	if err != nil {
		slog.Error("Error getting environment", "error", err)
		os.Exit(1)
	}
	spew.Dump(&env.JSON200)
}
