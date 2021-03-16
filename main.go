package main

import (
	"net/http"

	"github.com/bczarnobay/go-matrix/routes"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	routes.Setup()

	http.ListenAndServe(":8080", nil)
}
