package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt)
    <-c
	fmt.Println("Shutting down")
}
