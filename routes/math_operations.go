package routes

import (
	"fmt"
	"net/http"

	"github.com/bczarnobay/go-matrix/services"
)

func sumController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}

	result := services.Sum(records)

	response := fmt.Sprint(result)
	fmt.Fprint(w, response)
}

func multiplyController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}

	result := services.Multiply(records)

	response := fmt.Sprint(result)
	fmt.Fprint(w, response)
}
