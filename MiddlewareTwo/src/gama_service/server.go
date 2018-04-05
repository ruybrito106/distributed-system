package gama_service

import (
	newctx "context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/tracing/opentracing"
	tGrpc "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
)

var s = grpc.NewServer()

type executeGamaRequest struct{}
type executeGamaResponse struct{}

func Start(listener net.Listener, svc Service, logger log.Logger, tracer stdopentracing.Tracer) {

	opentracing.TraceServer(tracer, "ExecuteGama")(func(svc Service) endpoint.Endpoint {
		return func(ctx newctx.Context, req interface{}) (interface{}, error) {
			svc.ExecuteGama(ctx)
			return &executeGamaResponse{}, nil
		}
	}(svc))

	reflection.Register(s)

	level.Info(logger).Log("message", "Starting gama service gRPC server")

	if err := s.Serve(listener); err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to serve")
	}
}

type server struct {
	executeGama tGrpc.Handler
}
