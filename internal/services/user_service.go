package services

import "github.com/raphaelreis/goLabs/internal/models"

func GetUserList() []models.User {
	return []models.User{
		{ID: 1, Name: "Raphael", Email: "raphael@gmail.com"},
		{ID: 2, Name: "Rapha123", Email: "raphael123@gmail.com"},
	}
}
