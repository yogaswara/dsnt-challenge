package main

import (
	"fmt"
	"net/http"

	httphandler "dsnt-challenge/internal/adapters/handlers/http"
	"dsnt-challenge/internal/core/services"
	"dsnt-challenge/pkg/logger"
)

func main() {
	// Initialize logger
	logger.Init()

	// Initialize Services
	pingService := services.NewPingService()
	echoService := services.NewEchoService()

	// Initialize Handlers
	pingHandler := httphandler.NewPingHandler(pingService)
	echoHandler := httphandler.NewEchoHandler(echoService)

	// Setup Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler.Handle)
	mux.HandleFunc("/echo", echoHandler.Handle)

	// Start Server
	port := ":8080"
	logger.Info(fmt.Sprintf("Server is starting on port %s", port))
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Failed to start server", err)
	}
}
