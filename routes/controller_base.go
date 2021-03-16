package routes

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

func readInput(w http.ResponseWriter, r *http.Request) [][]string {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Method not allowed"}`))
		log.Println("Invalid method requested")
		return nil
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid input. Expected file."`))
		return nil
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil
	}
	if len(records) != len(records[0]) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid input. File should have same number of rows and columns."}`))
		return nil
	}
	return records
}
