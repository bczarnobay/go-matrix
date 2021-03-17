package services

import (
	"strconv"
)

func Invert(m [][]int) string {
	for r := 0; r < len(m); r++ {
		for c := 0; c < r; c++ {
			var tmp = m[r][c]
			m[r][c] = m[c][r]
			m[c][r] = tmp
		}
	}

	return Echo(m)
}

func Flatten(m [][]int) []string {
	var flat []string
	for _, row := range m {
		for _, column := range row {
			flat = append(flat, strconv.Itoa(column))
		}
	}
	return flat
}

func Echo(m [][]int) string {
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
	return response
}
