package operations

import (
	"context"

	"github.com/spf13/cobra"
)

type Resource struct {
	Kind     string
	Short    string
	Plural   string
	Singular string
	List     interface{}
	Get      interface{}
	Delete   interface{}
}

var resources = []*Resource{
	{
		Kind:     "Environment",
		Short:    "env",
		Plural:   "environments",
		Singular: "environment",
	},
}

type Operations struct {
	BaseURL string
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
			if resource.Plural == operation[1] {
				resource.Delete = fn
				break
			}
		}
		return
	}
	// fmt.Println("operation", operation)
	// if operation[0] == "u" {
	// 	applyRegister(ctx, operation, fn)
	// 	return
	// }

}

func (r *Operations) MainCommand() []*cobra.Command {
	return []*cobra.Command{
		// GetCmd,
		RemoveCmd,
		ApplyCmd,
		r.GetCmd(),
		r.LoginCmd(),
		r.LogoutCmd(),
		r.SetContextCmd(),
		r.GetContextCmd(),
	}
}

func (r *Operations) SetBaseURL(url string) {
	r.BaseURL = url
}
