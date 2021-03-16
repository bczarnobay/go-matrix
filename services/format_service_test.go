package services

import (
	"reflect"
	"testing"
)

func TestInvert(t *testing.T) {
	var input = [][]string{{"1", "2"}, {"3", "4"}}
	var expected = [][]string{{"1", "3"}, {"2", "4"}}

	Invert(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Matrix was not inverted properly")
	}
}

func TestFlatten(t *testing.T) {
	var input = [][]string{{"1", "2"}, {"3", "4"}}
	var expected = []string{"1", "2", "3", "4"}

	result := Flatten(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Matrix was not flatted properly")
	}
}
