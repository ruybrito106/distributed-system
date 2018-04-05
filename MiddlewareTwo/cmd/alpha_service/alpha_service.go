package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/ruybrito106/distributed-system/MiddlewareTwo/src/alpha_service"

	"github.com/streadway/amqp"
)

func main() {

	logger := log.NewNopLogger()
	tracer := stdopentracing.NoopTracer{}

	connPub, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		level.Error(logger).Log("err", err, "message", "Unable to stablish connection with MQ")
		return
	}

	defer connPub.Close()

	ch, err := connPub.Channel()
	if err != nil {
		level.Error(logger).Log("err", err, "message", "Unable to initialize channel")
		return
	}

	ch.QueueDeclare("queueA", false, false, false, false, nil)
	ch.QueueDeclare("queueB", false, false, false, false, nil)
	ch.QueueDeclare("queueG", false, false, false, false, nil)

	initialTemperature := 7.0 * rand.Float64()

	service := alpha_service.NewService(logger, ch, tracer, initialTemperature)
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		level.Error(logger).Log("err", err, "message", "Failed to listen")
		return
	}

	level.Info(logger).Log(
		"port", "5000",
		"message", "Starting alpha service at 5000",
	)

	go alpha_service.Start(listener, service, logger, tracer)

	go func() {
		for {
			service.ExecuteAlpha(context.Background())
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	connSub, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		level.Error(logger).Log("err", err, "message", "Unable to stablish connection with MQ")
		return
	}

	defer connSub.Close()

	chSub, err := connSub.Channel()
	if err != nil {
		level.Error(logger).Log("err", err, "message", "Unable to initialize channel")
		return
	}

	msgs, err := chSub.Consume(
		"queueA",
		"alpha",
		true,
		false,
		false,
		false,
		nil,
	)

	resolve := func(init string) string {
		switch init {
		case "A":
			return "Alpha"
		case "B":
			return "Beta"
		case "G":
			return "Gama"
		}
		return "Unknown"
	}

	go func() {
		for d := range msgs {
			msgString := string(d.Body)
			mapp := strings.Split(msgString, ":")

			if mapp[1] == "A" {
				fmt.Println(
					fmt.Sprintf(
						"%s service received value %s from %s",
						resolve(mapp[1]),
						mapp[2],
						resolve(mapp[0]),
					),
				)
			}
		}
	}()

	time.Sleep(time.Duration(100) * time.Second)

}
