package routes

import (
	"fmt"
	"net/http"

	"github.com/bczarnobay/go-matrix/services"
)

func echoController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}
	response, _ := services.Echo(records)
	fmt.Fprint(w, response)
}

func invertController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}
	response, _ := services.Invert(records)
	fmt.Fprint(w, response)
}

func flattenController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}
	fmt.Fprint(w, services.Flatten(records))
}
