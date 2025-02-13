package cli

import (
	"reflect"

	"github.com/beamlit/toolkit/sdk"
)

type Resource struct {
	Kind       string
	Short      string
	Plural     string
	Singular   string
	SpecType   reflect.Type
	List       interface{}
	Get        interface{}
	Delete     interface{}
	Put        interface{}
	Post       interface{}
	WithStatus bool
}

var resources = []*Resource{
	{
		Kind:     "Policy",
		Short:    "pol",
		Plural:   "policies",
		Singular: "policy",
		SpecType: reflect.TypeOf(sdk.Policy{}),
	},
	{
		Kind:       "Model",
		Short:      "ml",
		Plural:     "models",
		Singular:   "model",
		SpecType:   reflect.TypeOf(sdk.Model{}),
		WithStatus: true,
	},
	{
		Kind:       "Function",
		Short:      "fn",
		Plural:     "functions",
		Singular:   "function",
		SpecType:   reflect.TypeOf(sdk.Function{}),
		WithStatus: true,
	},
	{
		Kind:       "Agent",
		Short:      "ag",
		Plural:     "agents",
		Singular:   "agent",
		SpecType:   reflect.TypeOf(sdk.Agent{}),
		WithStatus: true,
	},
	{
		Kind:     "IntegrationConnection",
		Short:    "ic",
		Plural:   "integrationconnections",
		Singular: "integrationconnection",
		SpecType: reflect.TypeOf(sdk.IntegrationConnection{}),
	},
}
