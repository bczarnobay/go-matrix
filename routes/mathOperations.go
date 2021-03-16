package routes

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

func sumController(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var sum int
	var row []string

	for r := 0; r < len(records); r++ {
		row = records[r]
		for c := 0; c < len(row); c++ {
			num, _ := strconv.Atoi(row[c])
			sum = sum + num
		}
	}

	response := fmt.Sprint(sum)
	fmt.Fprint(w, response)
}

func multiplyController(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	result := 1
	var row []string

	for r := 0; r < len(records); r++ {
		row = records[r]
		for c := 0; c < len(row); c++ {
			num, _ := strconv.Atoi(row[c])
			result = result * num
		}
	}

	response := fmt.Sprint(result)
	fmt.Fprint(w, response)
}
