package beta_service

import (
	newctx "context"
	"net"

	"github.com/ruybrito106/MiddlewareOne/src/beta_service/proto"

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

type executeBetaRequest struct {
	Variable int
}

type executeBetaResponse struct{}

func Start(listener net.Listener, svc Service, logger log.Logger, tracer stdopentracing.Tracer) {

	var executeBeta endpoint.Endpoint
	executeBeta = opentracing.TraceServer(tracer, "ExecuteBeta")(func(svc Service) endpoint.Endpoint {
		return func(ctx newctx.Context, req interface{}) (interface{}, error) {
			request := req.(executeBetaRequest)
			svc.ExecuteBeta(ctx, request.Variable)
			return &executeBetaResponse{}, nil
		}
	}(svc))

	serv := makeGRPCServer(executeBeta, logger, tracer)
	proto.RegisterBetaServiceServer(s, serv)

	reflection.Register(s)

	level.Info(logger).Log("message", "Starting beta service gRPC server")

	if err := s.Serve(listener); err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to serve")
	}
}

type server struct {
	executeBeta tGrpc.Handler
}

func makeGRPCServer(e endpoint.Endpoint, logger log.Logger, tracer stdopentracing.Tracer) proto.BetaServiceServer {
	return server{
		executeBeta: tGrpc.NewServer(
			e,
			decodeExecuteBetaRequest,
			encodeExecuteBetaResponse,
			tGrpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "ExecuteBeta", logger),
			),
		),
	}
}

func (s server) ExecuteBeta(ctx newctx.Context, req *proto.ExecuteBetaRequest) (*proto.ExecuteBetaResponse, error) {
	_, rep, err := s.executeBeta.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.ExecuteBetaResponse), nil
}

func decodeExecuteBetaRequest(_ newctx.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.ExecuteBetaRequest)
	return executeBetaRequest{int(*req.Variable)}, nil
}

func encodeExecuteBetaResponse(_ newctx.Context, resp interface{}) (interface{}, error) {
	return &proto.ExecuteBetaResponse{nil}, nil
}
