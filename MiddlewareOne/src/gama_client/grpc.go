package gama_client

import (
	"context"
	"time"

	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/gama_service/proto"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	opentracing "github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type grpcGamaServiceClient struct {
	client proto.GamaServiceClient
}

func NewGrpcGamaServiceClient(addr string, tracer opentracing.Tracer) (GamaServiceClient, error) {
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

	var client GamaServiceClient
	client = &grpcGamaServiceClient{client: proto.NewGamaServiceClient(conn)}
	return client, nil
}

func (c *grpcGamaServiceClient) ExecuteGama(ctx context.Context, variable int) {

	newVar := int32(variable)

	req := proto.ExecuteGamaRequest{
		Variable: &newVar,
	}

	c.client.ExecuteGama(ctx, &req)

}
