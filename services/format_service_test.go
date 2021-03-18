package services

import (
	"errors"
	"testing"
)

type testCase struct {
	name     string
	input    [][]int
	expected string
	err      string
}

func TestInvert(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,3\n2,4", ""},
		{"2x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}}, "", errors.New(NOT_QUADRATIC).Error()},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := Invert(test.input)

			if err != nil {
				if err.Error() != test.err {
					t.Errorf("Wrong error message.Expected \n%v\n and got: \n%v\n", test.err, err)
				}
			}
			if result != test.expected {
				t.Errorf("Matrix was not inverted properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}

}

func TestFlatten(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,2,3,4", ""},
		{"2x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}}, "1,2,3,4,5,6", ""},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Flatten(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not inverted properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}

func TestEcho(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,2\n3,4", ""},
		{"2x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}}, "", errors.New(NOT_QUADRATIC).Error()},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := Echo(test.input)

			if err != nil {
				if err.Error() != test.err {
					t.Errorf("Wrong error message.Expected \n%v\n and got: \n%v\n", test.err, err)
				}
			}
			if result != test.expected {
				t.Errorf("Matrix was not echoed properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}
