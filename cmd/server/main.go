package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/314159otr/Peril/internal/routing"
	"github.com/314159otr/Peril/internal/pubsub"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	const CONN_STRING string = "amqp://guest:guest@localhost:5672/"
	con, err := amqp.Dial(CONN_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	fmt.Println("Connection was successful")

	channel, err := con.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
		IsPaused: true,
	})

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt)
    <-c
	fmt.Println("Shutting down")
}
