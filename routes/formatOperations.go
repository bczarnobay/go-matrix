package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bczarnobay/go-matrix/services"
)

func echoController(w http.ResponseWriter, r *http.Request) {
	records := readInput(w, r)

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func invertController(w http.ResponseWriter, r *http.Request) {
	records := readInput(w, r)

	services.Invert(records)

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func flattenController(w http.ResponseWriter, r *http.Request) {
	records := readInput(w, r)

	flat := services.Flatten(records)

	response := fmt.Sprint(strings.Join(flat, ","))
	fmt.Fprint(w, response)
}
