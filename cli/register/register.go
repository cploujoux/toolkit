package register

import (
	"context"

	"github.com/spf13/cobra"
)

type Register interface {
	CliCommand(ctx context.Context, operationId string, fn interface{})
	MainCommand() []*cobra.Command
}
