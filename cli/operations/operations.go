package operations

import (
	"context"

	"github.com/spf13/cobra"
)

type Operations struct {
	BaseURL string
}

func (r *Operations) CliCommand(ctx context.Context, operationId string, fn interface{}) {
	operation := formatOperationId(operationId)
	if operation[0] == "list" {
		GetRegister(ctx, operation, fn)
		return
	}
	if operation[0] == "get" {
		GetRegister(ctx, operation, fn)
		return
	}
	if operation[0] == "delete" {
		RemoveRegister(ctx, operation, fn)
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
		GetCmd,
		RemoveCmd,
		ApplyCmd,
		r.LoginCmd(),
		r.LogoutCmd(),
		r.SetContextCmd(),
		r.GetContextCmd(),
	}
}

func (r *Operations) SetBaseURL(url string) {
	r.BaseURL = url
}
