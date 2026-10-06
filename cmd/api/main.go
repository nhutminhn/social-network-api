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
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found, using default environment variables")
	}

	// Get port from the .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback port if PORT is not set
	}

	// Format server address
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
		log.Fatalf("Error running server: %v", err)
	}
}
