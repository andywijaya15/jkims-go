package config

import (
	"os"

	"github.com/andywijaya15/jkims-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IntiDB() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.People{})
}
