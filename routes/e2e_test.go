package routes

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type testInput struct {
	endpoint string
	value    string
	method   string
}

type testCase struct {
	name     string
	input    testInput
	expected string
}

func createTestRequest(input testInput) (*http.Request, error) {
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	csv, err := w.CreateFormFile("file", "matrix.csv")
	if err != nil {
		return nil, errors.New("There was an error creating the input file")
	}

	_, err = io.Copy(csv, strings.NewReader(input.value))
	if err != nil {
		return nil, errors.New("There was an error with the input string")
	}

	w.Close()

	testRequest, err := http.NewRequest(input.method, input.endpoint, body)
	if err != nil {
		return nil, errors.New("There was an error creating the test request")
	}

	testRequest.Header.Add("Content-Type", w.FormDataContentType())
	return testRequest, nil
}

func TestValidationsUsingEcho(t *testing.T) {
	var testCases = []testCase{
		{"4x4 matrix", testInput{"/echo", "1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16", "POST"}, "1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16"},
		{"4x4 matrix with GET", testInput{"/echo", "1,2,3,4\n5,6,7,8\n9,10,11,12", "GET"}, errors.New(NOT_ALLOWED).Error()},
		{"2x3 matrix", testInput{"/echo", "1,2\n3,4\n5,6", "POST"}, errors.New(NOT_QUADRATIC).Error()},
		{"1x1 matrix", testInput{"/echo", "1", "POST"}, errors.New(INVALID_SIZE).Error()},
		{"3x3 matrix with invalid character", testInput{"/echo", "1,2,3\n4,5,6\n7,8,a", "POST"}, errors.New(INVALID_CHARACTERS).Error()},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			request, err := createTestRequest(test.input)
			if err != nil {
				t.Errorf("Ops! Something happened while trying to create test request: %e", err)
			}

			rec := httptest.NewRecorder()
			controller := http.HandlerFunc(echoController)

			controller.ServeHTTP(rec, request)
			actual := rec.Body.String()

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected \n%v\n and got: \n%v\n", test.expected, actual)
			}
		})
	}
}

func TestSum(t *testing.T) {
	var testCases = []testCase{
		{"4x4 matrix", testInput{"/sum", "1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16", "POST"}, "136"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			request, err := createTestRequest(test.input)
			if err != nil {
				t.Errorf("Ops! Something happened while trying to create test request: %e", err)
			}

			rec := httptest.NewRecorder()
			controller := http.HandlerFunc(sumController)

			controller.ServeHTTP(rec, request)
			actual := rec.Body.String()

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected \n%v\n and got: \n%v\n", test.expected, actual)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	var testCases = []testCase{
		{"4x4 matrix", testInput{"/sum", "1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16", "POST"}, "20922789888000"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			request, err := createTestRequest(test.input)
			if err != nil {
				t.Errorf("Ops! Something happened while trying to create test request: %e", err)
			}

			rec := httptest.NewRecorder()
			controller := http.HandlerFunc(multiplyController)

			controller.ServeHTTP(rec, request)
			actual := rec.Body.String()

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected \n%v\n and got: \n%v\n", test.expected, actual)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	var testCases = []testCase{
		{"2x2 matrix", testInput{"/flatten", "1,2,3\n4,5,6\n7,8,9", "POST"}, "1,2,3,4,5,6,7,8,9"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			request, err := createTestRequest(test.input)
			if err != nil {
				t.Errorf("Ops! Something happened while trying to create test request: %e", err)
			}

			rec := httptest.NewRecorder()
			controller := http.HandlerFunc(flattenController)

			controller.ServeHTTP(rec, request)
			actual := rec.Body.String()

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected \n%v\n and got: \n%v\n", test.expected, actual)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	var testCases = []testCase{
		{"2x2 matrix", testInput{"/flatten", "1,2,3\n4,5,6\n7,8,9", "POST"}, "1,4,7\n2,5,8\n3,6,9"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			request, err := createTestRequest(test.input)
			if err != nil {
				t.Errorf("Ops! Something happened while trying to create test request: %e", err)
			}

			rec := httptest.NewRecorder()
			controller := http.HandlerFunc(invertController)

			controller.ServeHTTP(rec, request)
			actual := rec.Body.String()

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected \n%v\n and got: \n%v\n", test.expected, actual)
			}
		})
	}
}
