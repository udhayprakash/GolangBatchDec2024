package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}

	if Add(-1, 1) != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

// custom assertions

func AssertEqual(t *testing.T, expected, actual int, message string) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", message, expected, actual)
	}
}

func TestAdd3(t *testing.T) {
	result := Add(2, 3)
	AssertEqual(t, 5, result, "2 + 3 should equal 5")

	result = Add(-1, 1)
	AssertEqual(t, 0, result, "-1 + 1 should equal 0")
}
