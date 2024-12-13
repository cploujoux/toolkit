package cli

import (
	"context"
	"strings"
)

type Operations struct {
	BaseURL     string
	RunURL      string
	AppURL      string
	RegistryURL string
}

func (r *Operations) GetRegistryURL() string {
	registryURL := strings.Replace(r.RegistryURL, "https://", "", 1)
	registryURL = strings.Replace(registryURL, "http://", "", 1)
	return registryURL
}

func (r *Operations) CliCommand(ctx context.Context, operationId string, fn interface{}) {
	operation := formatOperationId(operationId)
	if operation[0] == "list" {
		for _, resource := range resources {
			if resource.Plural == operation[1] {
				resource.List = fn
				break
			}
		}
		return
	}
	if operation[0] == "get" {
		for _, resource := range resources {
			if resource.Singular == operation[1] {
				resource.Get = fn
				break
			}
		}
		return
	}
	if operation[0] == "delete" {
		for _, resource := range resources {
			if resource.Singular == operation[1] {
				resource.Delete = fn
				break
			}
		}
		return
	}
	if operation[0] == "put" || operation[0] == "update" {
		for _, resource := range resources {
			if resource.Singular == operation[1] {
				resource.Put = fn
				break
			}
		}
		return
	}
	if operation[0] == "create" {
		for _, resource := range resources {
			if resource.Singular == operation[1] {
				resource.Post = fn
				break
			}
		}
		return
	}

}

func (r *Operations) SetBaseURL(url string) {
	r.BaseURL = url
}
