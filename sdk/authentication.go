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
