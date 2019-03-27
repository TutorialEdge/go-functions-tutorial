package main

import "testing"

func TestSimpleFunction(t *testing.T) {
	expected := "Elliot Forbes"
	result := SimpleFunction("Elliot", "Forbes")

	if expected != result {
		t.Error("The result does not match what we expected")
	}
}

func TestSimpleFuncWithError(t *testing.T) {
	expected := "Elliot Forbes"
	result, err := SimpleFuncWithError("Elliot", "Forbes")
	if err != nil {
		t.Error("Our Function Threw an Error")
	}
	if expected != result {
		t.Error("The result does not match what we expected")
	}
}

func TestNegativeSimpleFuncWithError(t *testing.T) {
	result, err := SimpleFuncWithError("E", "Forbes")
	if err == nil {
		t.Error("Our function never threw an error")
	}
	if result != "" {
		t.Error("Expected an empty string returned")
	}
}

func TestChallenge01(t *testing.T) {
	var tests = []struct {
		value1   int
		value2   int
		expected int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 2, 2},
		{-5, -3, -8},
		{99999, 1, 100000},
	}

	for _, test := range tests {
		result := Add(test.value1, test.value2)
		if result != test.expected {
			t.Error("Expected output did not match the result")
		}
	}
}
