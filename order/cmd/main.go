package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

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
	defer dbAdapter.Close()

	portStr := os.Getenv("ORDER_SERVICE_PORT")
	if portStr == "" {
		log.Fatal("port is missing, please provide port in env")
	}

	paymentAdapter, err := payment.NewAdapter(os.Getenv("PAYMENT_SERVICE_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	defer paymentAdapter.Close()

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("%d is invalid port, port must be a number", port)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, port)

	go func() {

		if err = grpcAdapter.Run(); err != nil {
			log.Fatal(err)
		}

	}()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("place order from http"))
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT")),
		Handler: mux,
	}

	go func() {

		log.Println("http server running on port", server.Addr)
		if err = server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

}
