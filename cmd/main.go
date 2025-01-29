package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/raphaelreis/goLabs/internal/api"
	"github.com/raphaelreis/goLabs/internal/utils"
)

func main() {
	fmt.Println("Hello, World!")

	// Teste de UUID
	id := uuid.New()
	fmt.Println("Generated UUID:", id)

	// Chamar função do pacote interno
	utils.PrintMessage("Executando projeto robusto com Go!")

	// Iniciar servidor HTTP
	r := api.SetupRouter()
	r.Run(":8080")
}
