package alpha_client

import (
	"context"
	"time"

	"github.com/ruybrito106/MiddlewareOne/src/alpha_service/proto"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	opentracing "github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type grpcAlphaServiceClient struct {
	client proto.AlphaServiceClient
}

func NewGrpcAlphaServiceClient(addr string, tracer opentracing.Tracer) (AlphaServiceClient, error) {
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

	var client AlphaServiceClient
	client = &grpcAlphaServiceClient{client: proto.NewAlphaServiceClient(conn)}
	return client, nil
}

func (c *grpcAlphaServiceClient) ExecuteAlpha(ctx context.Context, variable int) {

	newVar := int32(variable)

	req := proto.ExecuteAlphaRequest{
		Variable: &newVar,
	}

	c.client.ExecuteAlpha(ctx, &req)

}
