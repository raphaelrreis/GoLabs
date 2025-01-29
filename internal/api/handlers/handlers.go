package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoLabs API!"})
}

func GetUsers(c *gin.Context) {
	users := []string{"Alice", "Bob", "Charlie"}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
