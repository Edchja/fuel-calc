package logic

import (
	"math"
	"testing"
)

func TestCalculateKilometerPrice_ValidInput(t *testing.T) {
	fuelConsumption := 17.0
	fuelPrice := 2.02
	expected := 0.34

	result, err := CalculateKilometerPrice(fuelConsumption, fuelPrice)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if roundFloat(result) != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}
}

func TestCalculateKilometerPrice_ZeroValues(t *testing.T) {
	zeroValue := 0.0
	result, err := CalculateKilometerPrice(zeroValue, zeroValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if roundFloat(result) != zeroValue {
		t.Errorf("expected %.2f, got %.2f", zeroValue, result)
	}
}

func TestParseToFloat_ValidInput(t *testing.T) {
	input := "8.5"
	expected := 8.5

	result, err := ParseToFloat(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestParseToFloat_EmptyInput(t *testing.T) {
	input := ""
	expectedErr := "input value can't be empty"

	_, err := ParseToFloat(input)
	if err == nil || err.Error() != expectedErr {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
}

func TestParseToFloat_InvalidInput(t *testing.T) {
	input := "abc"
	expectedErr := "invalid input value"

	_, err := ParseToFloat(input)
	if err == nil || err.Error() != expectedErr {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
}

func TestParseToFloat_NegativeInput(t *testing.T) {
	input := "-5.0"
	expectedErr := "input value can't be negative"

	_, err := ParseToFloat(input)
	if err == nil || err.Error() != expectedErr {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
}

func roundFloat(result float64) float64 {
	return math.Round(result*100) / 100
}
