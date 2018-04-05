package beta_service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

var (
	resolve = func(init string) string {
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
)

type temperatureControlMiddleware struct {
	next    Service
	channel *amqp.Channel
}

func (t temperatureControlMiddleware) ExecuteBeta(ctx context.Context) float64 {

	clients := []string{"A", "G"}
	chosenDestiny := clients[rand.Intn(len(clients))]
	sender := "B"

	nextMeasure := t.next.ExecuteBeta(ctx)
	formattedMsg := strconv.FormatFloat(nextMeasure, 'f', 4, 64)

	msg := strings.Join(
		[]string{
			sender,
			chosenDestiny,
			formattedMsg,
		},
		":",
	)

	t.channel.Publish(
		"",
		"queue"+chosenDestiny,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	fmt.Println(
		fmt.Sprintf(
			"%s service sent value %s to service %s",
			resolve(sender),
			formattedMsg,
			resolve(chosenDestiny),
		),
	)

	return nextMeasure

}
