package cli

import (
	"github.com/tmp-moon/toolkit/sdk"
)

func getAuthProvider() sdk.AuthProvider {
	workspaceName := sdk.CurrentContext()
	credentials := sdk.LoadCredentials(workspaceName)
	var provider sdk.AuthProvider
	if credentials.AccessToken != "" {
		provider = sdk.NewBearerTokenProvider(credentials.AccessToken, workspaceName)
	} else if credentials.APIKey != "" {
		provider = sdk.NewApiKeyProvider(credentials.APIKey, workspaceName)
	} else {
		provider = sdk.NewPublicProvider()
	}

	return provider
}
