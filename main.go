package main

import (
	"net/http"
	"runtime"

	"github.com/bczarnobay/go-matrix/routes"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	routes.Setup()

	http.ListenAndServe(":8080", nil)
}
