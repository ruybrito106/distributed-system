package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/beta_client"
	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/gama_client"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/alpha_service"
)

func main() {

	logger := log.NewNopLogger()
	tracer := stdopentracing.NoopTracer{}

	clients := alpha_service.Clients{}
	clients.BetaClient, _ = beta_client.NewGrpcBetaServiceClient("localhost:5002", nil)
	clients.GamaClient, _ = gama_client.NewGrpcGamaServiceClient("localhost:5003", nil)

	service := alpha_service.NewService(clients, logger, tracer)
	port := "5001"
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to listen")
		return
	}

	level.Info(logger).Log(
		"port", port,
		"message", fmt.Sprintf("Starting alpha service at %s", port),
	)

	go alpha_service.Start(listener, service, logger, tracer)

	service.ExecuteAlpha(context.Background(), 10)

	time.Sleep(time.Duration(100) * time.Second)

}
