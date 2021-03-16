package routes

import (
	"fmt"
	"net/http"

	"github.com/bczarnobay/go-matrix/services"
)

func sumController(w http.ResponseWriter, r *http.Request) {
	records := readInput(w, r)

	result := services.Sum(records)

	response := fmt.Sprint(result)
	fmt.Fprint(w, response)
}

func multiplyController(w http.ResponseWriter, r *http.Request) {
	records := readInput(w, r)

	result := services.Multiply(records)

	response := fmt.Sprint(result)
	fmt.Fprint(w, response)
}
