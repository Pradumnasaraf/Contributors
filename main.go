package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/pradumnasaraf/go-api/router"
)

func main(){
	r := router.Router()
	fmt.Println("Server is started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Print("Listening at port: 8080")
}