package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelreis/goLabs/internal/api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.HomeHandler)
	r.GET("/users", handlers.GetUsers)
	return r
}
