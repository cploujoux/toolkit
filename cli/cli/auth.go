package cli

import (
	"github.com/tmp-moon/toolkit/sdk"
)

func getAuthProvider() sdk.AuthProvider {
	workspace := sdk.CurrentContext()
	credentials := sdk.LoadCredentials(workspace)

	var provider sdk.AuthProvider
	if credentials.AccessToken != "" {
		provider = sdk.NewBearerTokenProvider(credentials.AccessToken, workspace)
	} else if credentials.APIKey != "" {
		provider = sdk.NewApiKeyProvider(credentials.APIKey, workspace)
	} else {
		provider = sdk.NewPublicProvider()
	}

	return provider
}
