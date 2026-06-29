package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/314159otr/Peril/internal/routing"
	"github.com/314159otr/Peril/internal/gamelogic"
	"github.com/314159otr/Peril/internal/pubsub"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	const CONN_STRING string = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(CONN_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Client Connection was successful")
	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatal(err)
	}

	_, queue, err := pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilDirect,
		routing.PauseKey + "." + username,
		routing.PauseKey,
		pubsub.SimpleQueueTransient,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Queue %v declared and bound!\n", queue.Name)

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt)
    <-c
	fmt.Println("Shutting down")
}
