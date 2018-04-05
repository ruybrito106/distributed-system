package alpha_service

import (
	"context"
	"math/rand"

	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/streadway/amqp"
)

type Service interface {
	ExecuteAlpha(context.Context) float64
}

type basicService struct {
	logger      log.Logger
	channel     *amqp.Channel
	tracer      opentracing.Tracer
	temperature float64
}

func NewService(logger log.Logger, channel *amqp.Channel, tracer opentracing.Tracer, temperature float64) Service {
	var service Service
	service = basicService{logger, channel, tracer, temperature}
	service = temperatureControlMiddleware{service, channel}
	return service
}

func (s basicService) ExecuteAlpha(ctx context.Context) float64 {
	variation := rand.Float64()
	s.temperature += variation
	return s.temperature
}

func (s basicService) ResetTemperature() {
	s.temperature = 7.0 * rand.Float64()
}
