package sdk

import (
	"context"
	"net/http"
)

type PublicProvider struct{}

func NewPublicProvider() *PublicProvider {
	return &PublicProvider{}
}

func (s *PublicProvider) Intercept(ctx context.Context, req *http.Request) error {
	return nil
}

type AuthProvider interface {
	Intercept(ctx context.Context, req *http.Request) error
}

type RunClientWithCredentials struct {
	ApiURL      string
	RunURL      string
	Credentials Credentials
	Workspace   string
}

func NewClientWithCredentials(config RunClientWithCredentials) (*ClientWithResponses, error) {
	var provider AuthProvider
	if config.Credentials.APIKey != "" {
		provider = NewApiKeyProvider(config.Credentials, config.Workspace)
	} else if config.Credentials.AccessToken != "" {
		provider = NewBearerTokenProvider(config.Credentials, config.Workspace, config.ApiURL)
	} else {
		provider = NewPublicProvider()
	}
	return NewClientWithResponses(config.ApiURL, config.RunURL, WithRequestEditorFn(provider.Intercept))
}
