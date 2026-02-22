package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/skyespirates/microservices/payment/config"
	"github.com/skyespirates/microservices/payment/internal/adapters/db"
	"github.com/skyespirates/microservices/payment/internal/adapters/grpc"
	"github.com/skyespirates/microservices/payment/internal/application/core/api"
)

func main() {

	godotenv.Load()

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
