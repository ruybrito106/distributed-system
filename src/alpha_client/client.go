package alpha_client

import (
	"context"
)

type AlphaServiceClient interface {
	ExecuteAlpha(context.Context, int)
}
