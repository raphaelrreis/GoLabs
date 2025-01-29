package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Not Connected DB:", err)
	}
	return db
}
