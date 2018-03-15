package beta_service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ruybrito106/MiddlewareOne/src/alpha_client"
	"github.com/ruybrito106/MiddlewareOne/src/gama_client"
)

type Service interface {
	ExecuteBeta(context.Context, int)
}

type basicService struct {
	logger      log.Logger
	tracer      opentracing.Tracer
	alphaClient alpha_client.AlphaServiceClient
	gamaClient  gama_client.GamaServiceClient
}

type Clients struct {
	AlphaClient alpha_client.AlphaServiceClient
	GamaClient  gama_client.GamaServiceClient
}

func NewService(clients Clients, logger log.Logger, tracer opentracing.Tracer) Service {
	var service Service
	service = basicService{logger, tracer, clients.AlphaClient, clients.GamaClient}
	return service
}

func (s basicService) ExecuteBeta(ctx context.Context, variable int) {

	fmt.Println("Beta received value", variable)
	time.Sleep(time.Second)

	variable = variable + 1
	fmt.Println("Beta processed variable")
	time.Sleep(time.Second)

	switch rand.Intn(2) {
	case 0:
		fmt.Println("Beta sent value", variable, "to Alpha service")
		s.alphaClient.ExecuteAlpha(ctx, variable)
	case 1:
		fmt.Println("Beta sent value", variable, "to Gama service")
		s.gamaClient.ExecuteGama(ctx, variable)
	}

	time.Sleep(time.Second)

}
