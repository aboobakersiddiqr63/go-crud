package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/aboobakersiddiqr63/go-crud/routes"
)

func main() {

	fmt.Println("Go lang-API")

	r := router.Router()
	fmt.Println("Starting the server on port 4000")
	fmt.Println("debug error one")

	log.Fatal(http.ListenAndServe(":4000", r))
}
