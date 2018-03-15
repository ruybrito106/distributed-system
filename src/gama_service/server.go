package gama_service

import (
	newctx "context"
	"net"

	"github.com/ruybrito106/MiddlewareOne/src/gama_service/proto"

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

type executeGamaRequest struct {
	Variable int
}

type executeGamaResponse struct{}

func Start(listener net.Listener, svc Service, logger log.Logger, tracer stdopentracing.Tracer) {

	var executeGama endpoint.Endpoint
	executeGama = opentracing.TraceServer(tracer, "ExecuteGama")(func(svc Service) endpoint.Endpoint {
		return func(ctx newctx.Context, req interface{}) (interface{}, error) {
			request := req.(executeGamaRequest)
			svc.ExecuteGama(ctx, request.Variable)
			return &executeGamaResponse{}, nil
		}
	}(svc))

	serv := makeGRPCServer(executeGama, logger, tracer)
	proto.RegisterGamaServiceServer(s, serv)

	reflection.Register(s)

	level.Info(logger).Log("message", "Starting gama service gRPC server")

	if err := s.Serve(listener); err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to serve")
	}
}

type server struct {
	executeGama tGrpc.Handler
}

func makeGRPCServer(e endpoint.Endpoint, logger log.Logger, tracer stdopentracing.Tracer) proto.GamaServiceServer {
	return server{
		executeGama: tGrpc.NewServer(
			e,
			decodeExecuteGamaRequest,
			encodeExecuteGamaResponse,
			tGrpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "ExecuteGama", logger),
			),
		),
	}
}

func (s server) ExecuteGama(ctx newctx.Context, req *proto.ExecuteGamaRequest) (*proto.ExecuteGamaResponse, error) {
	_, rep, err := s.executeGama.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.ExecuteGamaResponse), nil
}

func decodeExecuteGamaRequest(_ newctx.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.ExecuteGamaRequest)
	return executeGamaRequest{int(*req.Variable)}, nil
}

func encodeExecuteGamaResponse(_ newctx.Context, resp interface{}) (interface{}, error) {
	return &proto.ExecuteGamaResponse{nil}, nil
}
