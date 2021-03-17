package routes

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
)

func readInput(w http.ResponseWriter, r *http.Request) [][]string {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%v", errors.New(NOT_ALLOWED))))
		return nil
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", errors.New(INVALID_INPUT_TYPE))))
		return nil
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err.Error())))
		return nil
	}

	if len(records) <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", errors.New(INVALID_SIZE))))
		return nil
	}

	if len(records) != len(records[0]) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", errors.New(NOT_QUADRATIC))))
		return nil
	}

	// intMatrix, err := convertMatrixToInt(records)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(fmt.Sprintf(`{"message":  "%s"}`, err.Error())))
	// }

	return records
}

// func convertMatrixToInt(m [][]string) ([][]int, error) {
// 	var converted [][]int
// 	for r := 0; r < len(m); r++ {
// 		for c := 0; c < r; c++ {
// 			num, err := strconv.Atoi(m[r][c])
// 			if err != nil {
// 				return nil, errors.New(INVALID_CHARACTERS)
// 			}
// 			converted[r][c] = num
// 		}
// 	}
// 	return converted, nil
// }
