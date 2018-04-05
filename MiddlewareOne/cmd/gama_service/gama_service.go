package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/alpha_client"
	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/beta_client"
	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/gama_service"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	stdopentracing "github.com/opentracing/opentracing-go"
)

func main() {

	logger := log.NewNopLogger()
	tracer := stdopentracing.NoopTracer{}

	clients := gama_service.Clients{}
	clients.AlphaClient, _ = alpha_client.NewGrpcAlphaServiceClient("localhost:5001", nil)
	clients.BetaClient, _ = beta_client.NewGrpcBetaServiceClient("localhost:5002", nil)

	service := gama_service.NewService(clients, logger, tracer)
	port := "5003"
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to listen")
		return
	}

	level.Info(logger).Log(
		"port", port,
		"message", fmt.Sprintf("Starting gama service at %s", port),
	)

	go gama_service.Start(listener, service, logger, tracer)

	time.Sleep(time.Duration(100) * time.Second)

}
