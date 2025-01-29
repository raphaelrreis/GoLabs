package server

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelreis/goLabs/internal/api"
)

func StartServer() error {
	r := gin.Default()

	// Configurar rotas
	api.SetupRouter(r)

	// Iniciar servidor na porta 8080
	return r.Run(":8080")
}
