package alpha_service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/beta_client"
	"github.com/ruybrito106/distributed-system/MiddlewareOne/src/gama_client"

	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
)

type Service interface {
	ExecuteAlpha(context.Context, int)
}

type basicService struct {
	logger     log.Logger
	tracer     opentracing.Tracer
	betaClient beta_client.BetaServiceClient
	gamaClient gama_client.GamaServiceClient
}

type Clients struct {
	BetaClient beta_client.BetaServiceClient
	GamaClient gama_client.GamaServiceClient
}

func NewService(clients Clients, logger log.Logger, tracer opentracing.Tracer) Service {
	var service Service
	service = basicService{logger, tracer, clients.BetaClient, clients.GamaClient}
	return service
}

func (s basicService) ExecuteAlpha(ctx context.Context, variable int) {

	fmt.Println("Alpha received value", variable)
	time.Sleep(time.Second)

	variable = variable + 1
	fmt.Println("Alpha processed variable")
	time.Sleep(time.Second)

	switch rand.Intn(2) {
	case 0:
		fmt.Println("Alpha sent value", variable, "to Beta")
		s.betaClient.ExecuteBeta(ctx, variable)
	case 1:
		fmt.Println("Alpha sent value", variable, "to Gama")
		s.gamaClient.ExecuteGama(ctx, variable)
	}

	time.Sleep(time.Second)

}
