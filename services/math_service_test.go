package services

import (
	"testing"
)

type testCaseMath struct {
	name     string
	input    [][]int
	expected int
}

func TestSum(t *testing.T) {
	var testCases = []testCaseMath{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, 10},
		{"3x3 Matrix with negative number", [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, 0}}, -36},
		{"2x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}}, 21},
		{"0x0 Matrix", [][]int{{}}, 0},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Sum(test.input)
			if result != test.expected {
				t.Errorf("Sum was incorrect. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	var testCases = []testCaseMath{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, 24},
		{"3x3 Matrix with 0", [][]int{{0, 2, 3}, {5, 6, 7}, {8, 9, 1}}, 0},
		{"3x3 Matrix, 1 to 9", [][]int{{1, 2, 3}, {5, 6, 7}, {8, 9, 4}}, 362880},
		{"3x3 Matrix, -1 to -9", [][]int{{-9, -1, -2}, {-3, -4, -5}, {-6, -7, -8}}, -362880},
		{"2x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}}, 720},
		{"0x0 Matrix", [][]int{{}}, 0},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Multiply(test.input)

			if result != test.expected {
				t.Errorf("Multiplication was incorrect. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}
