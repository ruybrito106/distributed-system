package gama_client

import (
	"context"
)

type GamaServiceClient interface {
	ExecuteGama(context.Context, int)
}
