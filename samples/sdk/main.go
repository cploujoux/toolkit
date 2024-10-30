package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/davecgh/go-spew/spew"
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

	ctx := context.Background()

	provider := beamlit.NewApiKeyProvider("BL_96UX6OFWA6JG9IA1NDE7CCMED22U8Z2F", "chris")

	client, err := beamlit.NewClientWithResponses(
		BASE_URL,
		beamlit.WithRequestEditorFn(provider.Intercept),
	)
	if err != nil {
		slog.Error("Error creating client", "error", err)
		os.Exit(1)
	}

	env, err := client.CreateOrUpdateEnvironmentWithResponse(ctx, "test", beamlit.CreateOrUpdateEnvironmentJSONRequestBody{
		Name:        beamlit.BlString("test"),
		DisplayName: beamlit.BlString("Test"),
	})
	if err != nil {
		slog.Error("Error getting environment", "error", err)
		os.Exit(1)
	}
	spew.Dump(&env.JSON200)
}
