package main

import (
	"log"
	"net/http"

	"github.com/bczarnobay/go-matrix/routes"
)

func main() {
	routes.Setup()

	http.ListenAndServe(":8080", nil)
	log.Println("Server is up and running on port :8080")
}
