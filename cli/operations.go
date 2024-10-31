package cli

import (
	"context"
	"reflect"

	"github.com/tmp-moon/toolkit/sdk"
)

type Resource struct {
	Kind     string
	Short    string
	Plural   string
	Singular string
	SpecType reflect.Type
	List     interface{}
	Get      interface{}
	Delete   interface{}
	Put      interface{}
}

var resources = []*Resource{
	{
		Kind:     "Environment",
		Short:    "env",
		Plural:   "environments",
		Singular: "environment",
		SpecType: reflect.TypeOf(sdk.Environment{}),
	},
	{
		Kind:     "Policy",
		Short:    "pl",
		Plural:   "policies",
		Singular: "policy",
		SpecType: reflect.TypeOf(sdk.Policy{}),
	},
	{
		Kind:     "Model",
		Short:    "ml",
		Plural:   "models",
		Singular: "model",
		SpecType: reflect.TypeOf(sdk.Policy{}),
	},
	{
		Kind:     "ModelProvider",
		Short:    "mlp",
		Plural:   "modelproviders",
		Singular: "modelprovider",
		SpecType: reflect.TypeOf(sdk.ModelProvider{}),
	},
}

type Operations struct {
	BaseURL string
	RunURL  string
}

func (r *Operations) CliCommand(ctx context.Context, operationId string, fn interface{}) {
	operation := formatOperationId(operationId)
	if len(operation) > 2 {
		return
	}
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
	if operation[0] == "put" {
		for _, resource := range resources {
			if resource.Singular == operation[1] {
				resource.Put = fn
				break
			}
		}
		return
	}

}

func (r *Operations) SetBaseURL(url string) {
	r.BaseURL = url
}
