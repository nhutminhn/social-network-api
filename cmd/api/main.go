package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load evironment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("warning: No .env file found, using default environment variables")
	}

	// Get port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Format address server
	addr := fmt.Sprintf(":%s", port)

	// Configure server timeouts to prevent resource leaks
	server := &http.Server{
		Addr:         addr,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Server running on port: %s\n", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error running server: %v\n", err)
	}
}
