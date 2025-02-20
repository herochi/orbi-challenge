package main

import (
	"fmt"
	"log"

	"github.com/herochi/orbi/service-b/internal/application"
	"github.com/herochi/orbi/service-b/internal/infrastructure/grpc"
	"github.com/herochi/orbi/service-b/internal/infrastructure/http"
	"github.com/herochi/orbi/service-b/internal/infrastructure/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Printf("DEPLOYING SERVICE B 🚀 \n")

	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	log.Printf("service b Connected to RabbitMQ")
	defer conn.Close()

	publisher, _ := rabbitmq.NewPublisher(conn)
	userClient := http.NewUserClient()
	notifyService := application.NewNotifyService(*publisher, *userClient)

	go rabbitmq.StartConsumer(conn, "service-a-notifies")
	grpc.StartGRPCServer(notifyService)
}
