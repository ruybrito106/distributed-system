package gama_service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ruybrito106/MiddlewareOne/src/alpha_client"
	"github.com/ruybrito106/MiddlewareOne/src/beta_client"
)

type Service interface {
	ExecuteGama(context.Context, int)
}

type basicService struct {
	logger      log.Logger
	tracer      opentracing.Tracer
	alphaClient alpha_client.AlphaServiceClient
	betaClient  beta_client.BetaServiceClient
}

type Clients struct {
	AlphaClient alpha_client.AlphaServiceClient
	BetaClient  beta_client.BetaServiceClient
}

func NewService(clients Clients, logger log.Logger, tracer opentracing.Tracer) Service {
	var service Service
	service = basicService{logger, tracer, clients.AlphaClient, clients.BetaClient}
	return service
}

func (s basicService) ExecuteGama(ctx context.Context, variable int) {

	fmt.Println("Gama received value", variable)
	time.Sleep(time.Second)

	variable = variable + 1
	fmt.Println("Gama processed variable")
	time.Sleep(time.Second)

	switch rand.Intn(2) {
	case 0:
		fmt.Println("Gama sent value", variable, "to Alpha service")
		s.alphaClient.ExecuteAlpha(ctx, variable)
	case 1:
		fmt.Println("Gama sent value", variable, "to Beta service")
		s.betaClient.ExecuteBeta(ctx, variable)
	}

	time.Sleep(time.Second)

}
