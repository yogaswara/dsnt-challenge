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

	// Initialize Handlers
	pingHandler := httphandler.NewPingHandler(pingService)

	// Setup Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler.Handle)

	// Start Server
	port := ":8080"
	logger.Info(fmt.Sprintf("Server is starting on port %s", port))
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Failed to start server", err)
	}
}
