package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Invert(m [][]int) (string, error) {
	if len(m) != len(m[0]) {
		return "", errors.New(NOT_QUADRATIC)
	}
	if len(m) == 0 {
		return "", errors.New(INVALID_SIZE)
	}
	for r := 0; r < len(m); r++ {
		for c := 0; c < r; c++ {
			var tmp = m[r][c]
			m[r][c] = m[c][r]
			m[c][r] = tmp
		}
	}
	response, _ := Echo(m)
	return response, nil
}

func Flatten(m [][]int) string {
	var flat []string
	for _, row := range m {
		for _, column := range row {
			flat = append(flat, strconv.Itoa(column))
		}
	}
	response := fmt.Sprint(strings.Join(flat, ","))
	return response
}

func Echo(m [][]int) (string, error) {
	if len(m) != len(m[0]) {
		return "", errors.New(NOT_QUADRATIC)
	}
	if len(m) == 0 {
		return "", errors.New(INVALID_SIZE)
	}

	var response string
	rowLen := len(m)
	for r, row := range m {
		for c, column := range row {
			response += strconv.Itoa(column)
			if c+1 < rowLen {
				response += ","
			}
		}

		if r+1 != rowLen {
			response += "\n"
		}
	}
	return response, nil
}
