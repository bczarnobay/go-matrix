package services

import (
	"testing"
)

func TestSum(t *testing.T) {
	var input = [][]int{{1, 1}, {1, 1}}

	result := Sum(input)

	if result != 4 {
		t.Errorf("Sum was incorrect. Expected 4 and got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	var input = [][]int{{1, 2}, {3, 4}}

	result := Multiply(input)

	if result != 24 {
		t.Errorf("Sum was incorrect. Expected 24 and got %d", result)
	}

}
