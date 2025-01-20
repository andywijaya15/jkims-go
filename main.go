package main

import (
	"log"

	"github.com/andywijaya15/jkims-go/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.IntiDB()
}
