package alpha_service

import (
	newctx "context"
	"net"

	"github.com/ruybrito106/MiddlewareOne/src/alpha_service/proto"

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

type executeAlphaRequest struct {
	Variable int
}

type executeAlphaResponse struct{}

func Start(listener net.Listener, svc Service, logger log.Logger, tracer stdopentracing.Tracer) {

	var executeAlpha endpoint.Endpoint
	executeAlpha = opentracing.TraceServer(tracer, "ExecuteAlpha")(func(svc Service) endpoint.Endpoint {
		return func(ctx newctx.Context, req interface{}) (interface{}, error) {
			request := req.(executeAlphaRequest)
			svc.ExecuteAlpha(ctx, request.Variable)
			return &executeAlphaResponse{}, nil
		}
	}(svc))

	serv := makeGRPCServer(executeAlpha, logger, tracer)
	proto.RegisterAlphaServiceServer(s, serv)

	reflection.Register(s)

	level.Info(logger).Log("message", "Starting alpha service gRPC server")

	if err := s.Serve(listener); err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to serve")
	}
}

type server struct {
	executeAlpha tGrpc.Handler
}

func makeGRPCServer(e endpoint.Endpoint, logger log.Logger, tracer stdopentracing.Tracer) proto.AlphaServiceServer {
	return server{
		executeAlpha: tGrpc.NewServer(
			e,
			decodeExecuteAlphaRequest,
			encodeExecuteAlphaResponse,
			tGrpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "ExecuteAlpha", logger),
			),
		),
	}
}

func (s server) ExecuteAlpha(ctx newctx.Context, req *proto.ExecuteAlphaRequest) (*proto.ExecuteAlphaResponse, error) {
	_, rep, err := s.executeAlpha.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.ExecuteAlphaResponse), nil
}

func decodeExecuteAlphaRequest(_ newctx.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.ExecuteAlphaRequest)
	return executeAlphaRequest{int(*req.Variable)}, nil
}

func encodeExecuteAlphaResponse(_ newctx.Context, resp interface{}) (interface{}, error) {
	return &proto.ExecuteAlphaResponse{nil}, nil
}
