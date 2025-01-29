package main

import (
	"fmt"
	"log"

	"github.com/raphaelreis/goLabs/internal/db"
	"github.com/raphaelreis/goLabs/internal/server"
)

func main() {
	fmt.Println("Starting GoLabs API...")

	err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = server.StartServer()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
