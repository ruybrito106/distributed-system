package beta_client

import (
	"context"
)

type BetaServiceClient interface {
	ExecuteBeta(context.Context, int)
}
