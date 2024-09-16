package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Config() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Using default environment variables")
	} else {
		log.Print("Environment variables loaded successfully from .env file")
	}

}
