package main

import (
	"os"

	"github.com/pradumnasaraf/go-api/router"
)

func main() {
	router := router.Router()
	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
