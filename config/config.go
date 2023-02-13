package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Config() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Some error occured. Err:", err)
	}

}
