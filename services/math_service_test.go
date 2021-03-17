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
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Sum(test.input)

			if result != test.expected {
				t.Errorf("Sum was incorrect. Expected \n%v\n and got: \n%v\n", test.expected, test.input)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	var testCases = []testCaseMath{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, 24},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Multiply(test.input)

			if result != test.expected {
				t.Errorf("Sum was incorrect. Expected \n%v\n and got: \n%v\n", test.expected, test.input)
			}
		})
	}
}
