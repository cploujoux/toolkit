package cli

import (
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
	{
		Kind:     "Location",
		Short:    "loc",
		Plural:   "locations",
		Singular: "location",
		SpecType: reflect.TypeOf(sdk.Location{}),
	},
	{
		Kind:     "Tool",
		Short:    "tl",
		Plural:   "tools",
		Singular: "tool",
		SpecType: reflect.TypeOf(sdk.Tool{}),
	},
}
