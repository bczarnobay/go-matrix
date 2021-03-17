package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bczarnobay/go-matrix/services"
)

func echoController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}

	fmt.Fprint(w, services.Echo(records))
}

func invertController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}

	fmt.Fprint(w, services.Invert(records))
}

func flattenController(w http.ResponseWriter, r *http.Request) {
	records, err := readInput(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return
	}
	flat := services.Flatten(records)

	response := fmt.Sprint(strings.Join(flat, ","))
	fmt.Fprint(w, response)
}
