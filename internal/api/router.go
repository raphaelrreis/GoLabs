package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelreis/goLabs/internal/api/handlers"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handlers.HomeHandler)
	r.GET("/users", handlers.GetUsers)
}
