package services

import (
	"strconv"
)

func Sum(m [][]string) int {
	var row []string
	var sum int

	for r := 0; r < len(m); r++ {
		row = m[r]
		for c := 0; c < len(row); c++ {
			num, _ := strconv.Atoi(row[c])
			sum = sum + num
		}
	}
	return sum
}

func Multiply(m [][]string) int {
	result := 1
	var row []string

	for r := 0; r < len(m); r++ {
		row = m[r]
		for c := 0; c < len(row); c++ {
			num, _ := strconv.Atoi(row[c])
			result = result * num
		}
	}

	return result
}
