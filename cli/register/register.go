package register

import (
	"context"
)

type Register interface {
	CliCommand(ctx context.Context, operationId string, fn interface{})
}
