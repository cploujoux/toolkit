package cli

import (
	"fmt"

	"github.com/tmp-moon/toolkit/sdk"
)

func getAuthProvider(workspaceName string) sdk.AuthProvider {
	if workspaceName == "" {
		workspaceName = sdk.CurrentContext()
	}
	credentials := sdk.LoadCredentials(workspaceName)
	fmt.Println("Workspaces:", workspaceName)
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
