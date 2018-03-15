package beta_client

import (
	"context"
	"time"

	"github.com/ruybrito106/MiddlewareOne/src/beta_service/proto"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	opentracing "github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type grpcBetaServiceClient struct {
	client proto.BetaServiceClient
}

func NewGrpcBetaServiceClient(addr string, tracer opentracing.Tracer) (BetaServiceClient, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinearWithJitter(20*time.Millisecond, 0.5)),
		grpc_retry.WithCodes(codes.Unavailable, codes.Aborted, codes.ResourceExhausted),
	}

	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)))

	if err != nil {
		return nil, err
	}

	var client BetaServiceClient
	client = &grpcBetaServiceClient{client: proto.NewBetaServiceClient(conn)}
	return client, nil
}

func (c *grpcBetaServiceClient) ExecuteBeta(ctx context.Context, variable int) {

	newVar := int32(variable)

	req := proto.ExecuteBetaRequest{
		Variable: &newVar,
	}

	c.client.ExecuteBeta(ctx, &req)

}
