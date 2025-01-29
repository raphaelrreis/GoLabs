package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bem Vindo ao GoLabs API!"})
}

func GetUsers(c *gin.Context) {
	users := []string{"Raphael", "Raphael123", "Pati"}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
