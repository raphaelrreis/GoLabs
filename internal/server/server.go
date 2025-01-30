package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raphaelreis/goLabs/internal/api"
	"github.com/raphaelreis/goLabs/internal/db"
)

func StartServer() {
	fmt.Println("🚀 Starting GoLabs API...")

	err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	r := gin.Default()
	api.SetupRouter(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	}
}
