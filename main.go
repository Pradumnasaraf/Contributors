package main

import (
	"os"

	"github.com/pradumnasaraf/go-api/router"
)

func main() {
	router := router.Router()
	router.Run(os.Getenv("API_HOST"))
}
