package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ruybrito106/MiddlewareOne/src/alpha_client"
	"github.com/ruybrito106/MiddlewareOne/src/beta_service"
	"github.com/ruybrito106/MiddlewareOne/src/gama_client"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	stdopentracing "github.com/opentracing/opentracing-go"
)

func main() {

	logger := log.NewNopLogger()
	tracer := stdopentracing.NoopTracer{}

	clients := beta_service.Clients{}
	clients.AlphaClient, _ = alpha_client.NewGrpcAlphaServiceClient("localhost:5001", nil)
	clients.GamaClient, _ = gama_client.NewGrpcGamaServiceClient("localhost:5003", nil)

	service := beta_service.NewService(clients, logger, tracer)
	port := "5002"
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to listen")
		return
	}

	level.Info(logger).Log(
		"port", port,
		"message", fmt.Sprintf("Starting beta service at %s", port),
	)

	go beta_service.Start(listener, service, logger, tracer)

	time.Sleep(time.Duration(100) * time.Second)

}
