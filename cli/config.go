package cli

import (
	"reflect"

	"github.com/beamlit/toolkit/sdk"
)

type Resource struct {
	Kind           string
	Short          string
	Plural         string
	Singular       string
	SpecType       reflect.Type
	ListParamsType reflect.Type
	GetParamsType  reflect.Type
	List           interface{}
	Get            interface{}
	Delete         interface{}
	Put            interface{}
	Post           interface{}
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
		Short:    "pol",
		Plural:   "policies",
		Singular: "policy",
		SpecType: reflect.TypeOf(sdk.Policy{}),
	},
	{
		Kind:           "Model",
		Short:          "ml",
		Plural:         "models",
		Singular:       "model",
		SpecType:       reflect.TypeOf(sdk.Model{}),
		ListParamsType: reflect.TypeOf(sdk.ListModelsParams{}),
		GetParamsType:  reflect.TypeOf(sdk.GetModelParams{}),
	},
	{
		Kind:     "ModelProvider",
		Short:    "mlp",
		Plural:   "modelproviders",
		Singular: "modelprovider",
		SpecType: reflect.TypeOf(sdk.ModelProvider{}),
	},
	{
		Kind:     "Location",
		Short:    "loc",
		Plural:   "locations",
		Singular: "location",
		SpecType: reflect.TypeOf(sdk.LocationResponse{}),
	},
	{
		Kind:           "Function",
		Short:          "fn",
		Plural:         "functions",
		Singular:       "function",
		SpecType:       reflect.TypeOf(sdk.Function{}),
		ListParamsType: reflect.TypeOf(sdk.ListFunctionsParams{}),
		GetParamsType:  reflect.TypeOf(sdk.GetFunctionParams{}),
	},
	{
		Kind:           "Agent",
		Short:          "ag",
		Plural:         "agents",
		Singular:       "agent",
		SpecType:       reflect.TypeOf(sdk.Agent{}),
		ListParamsType: reflect.TypeOf(sdk.ListAgentsParams{}),
		GetParamsType:  reflect.TypeOf(sdk.GetAgentParams{}),
	},
	{
		Kind:     "IntegrationConnection",
		Short:    "ic",
		Plural:   "integrationconnections",
		Singular: "integrationconnection",
		SpecType: reflect.TypeOf(sdk.IntegrationConnection{}),
	},
}
