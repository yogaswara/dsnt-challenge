package main

import (
	"fmt"
	"net/http"

	httphandler "dsnt-challenge/internal/adapters/handlers/http"
	"dsnt-challenge/internal/adapters/repository/memory"
	"dsnt-challenge/internal/core/services"
	"dsnt-challenge/pkg/logger"
)

func main() {
	// Initialize logger
	logger.Init()

	// Initialize Repositories
	bookRepo := memory.NewBookRepository()

	// Initialize Services
	pingService := services.NewPingService()
	echoService := services.NewEchoService()
	booksService := services.NewBooksService(bookRepo)
	authService := services.NewAuthService()

	// Initialize Handlers
	pingHandler := httphandler.NewPingHandler(pingService)
	echoHandler := httphandler.NewEchoHandler(echoService)
	booksHandler := httphandler.NewBooksHandler(booksService)
	authHandler := httphandler.NewAuthHandler(authService)

	// Setup Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler.Handle)
	mux.HandleFunc("/echo", echoHandler.Handle)
	mux.HandleFunc("/auth/token", authHandler.HandleToken)
	mux.HandleFunc("/books", httphandler.AuthMiddleware(authService, booksHandler.HandleBooks))
	mux.HandleFunc("/books/", httphandler.AuthMiddleware(authService, booksHandler.HandleBookByID))

	// Start Server
	port := ":8080"
	logger.Info(fmt.Sprintf("Server is starting on port %s", port))
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Failed to start server", err)
	}
}
