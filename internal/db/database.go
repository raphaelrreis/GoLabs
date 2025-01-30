package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "secret"),
		getEnv("DB_NAME", "golabs"),
		getEnv("DB_PORT", "5432"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = database
	fmt.Println("✅ Connected to database!")

	err = migrateDatabase()
	if err != nil {
		return err
	}

	return nil
}

func migrateDatabase() error {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	fmt.Println("Migrations applied successfully!")
	return nil
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique"`
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
