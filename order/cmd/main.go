package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/skyespirates/microservices/order/internal/adapters/db"
	"github.com/skyespirates/microservices/order/internal/adapters/grpc"
	"github.com/skyespirates/microservices/order/internal/adapters/payment"
	"github.com/skyespirates/microservices/order/internal/application/core/api"
)

func main() {

	godotenv.Load()

	dbAdapter, err := db.NewAdapter(os.Getenv("DATA_SOURCE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to database: error %v", err)
	}

	portStr := os.Getenv("APPLICATION_PORT")
	if portStr == "" {
		log.Fatal("port is missing, please provide port in env")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("%d is invalid port, port must be a number", port)
	}

	paymentAdapter, err := payment.NewAdapter(os.Getenv("PAYMENT_SERVICE_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, port)
	grpcAdapter.Run()

}
