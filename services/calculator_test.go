package services

import "testing"

func TestCalculatorAdd_Success(t *testing.T) {
	cal := &Calculator{
		Num1: 6,
		Num2: 1,
	}

	expected := 7
	actual, err := cal.Add()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *actual != float64(expected) {
		t.Errorf("Expected: %v", expected)
	}
}

func TestCalculatorAdd_Fail(t *testing.T) {
	cal := &Calculator{
		Num1: 6,
		Num2: 1,
	}

	expected := 5
	actual, err := cal.Add()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *actual == float64(expected) {
		t.Errorf("Expected: %v", expected)
	}
}

func TestCalculatorSub_Success(t *testing.T) {
	cal := &Calculator{
		Num1: 10,
		Num2: 5,
	}

	expected := 5
	actual, err := cal.Sub()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *actual != float64(expected) {
		t.Errorf("Expected: %v", expected)
	}
}

func TestCalculatorSub_Fail(t *testing.T) {
	cal := &Calculator{
		Num1: 6,
		Num2: 1,
	}

	expected := 10
	actual, err := cal.Sub()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *actual == float64(expected) {
		t.Errorf("Expected: %v", expected)
	}
}
