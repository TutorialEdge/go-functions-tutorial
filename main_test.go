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
