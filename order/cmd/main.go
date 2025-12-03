package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/skyespirates/microservices/order/internal/adapters/db"
	"github.com/skyespirates/microservices/order/internal/adapters/grpc"
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
		log.Fatal("port is missing")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("%d is invalid port", port)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, port)
	grpcAdapter.Run()
}
