package server

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelreis/goLabs/internal/api"
)

func StartServer() error {
	r := gin.Default()

	api.SetupRouter(r)

	return r.Run(":8080")
}
