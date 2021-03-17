package services

import (
	"testing"
)

type testCase struct {
	name     string
	input    [][]int
	expected string
}

func TestInvert(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,3\n2,4"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Invert(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not inverted properly. Expected \n%v\n and got: \n%v\n", test.expected, test.input)
			}
		})
	}

}

func TestFlatten(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,2,3,4"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Flatten(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not inverted properly. Expected \n%v\n and got: \n%v\n", test.expected, test.input)
			}
		})
	}
}

func TestEcho(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,2\n3,4"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Echo(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not echoed properly. Expected \n%v\n and got: \n%v\n", test.expected, test.input)
			}
		})
	}
}
