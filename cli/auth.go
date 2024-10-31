package cli

import (
	"github.com/tmp-moon/toolkit/sdk"
)

func getAuthProvider(workspaceName string) sdk.AuthProvider {
	if workspaceName == "" {
		workspaceName = sdk.CurrentWorkspace()
	}
	credentials := sdk.LoadCredentials(workspaceName)
	var provider sdk.AuthProvider
	if credentials.AccessToken != "" {
		provider = sdk.NewBearerTokenProvider(credentials, workspaceName)
	} else if credentials.APIKey != "" {
		provider = sdk.NewApiKeyProvider(credentials, workspaceName)
	} else {
		provider = sdk.NewPublicProvider()
	}

	return provider
}
