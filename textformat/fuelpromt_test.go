package textformat

import (
	"bufio"
	"errors"
	"strings"
	"testing"
)

type errorReader struct{}

func TestValidateAndFormatInput_ValidInputWithDot(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("123.45\n"))
	result, err := ValidateAndFormatInput()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "123.45" {
		t.Fatalf("expected '123.45', got %v", result)
	}
}

func TestValidateAndFormatInput_ValidInputWithComma(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("123,45\n"))
	result, err := ValidateAndFormatInput()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "123.45" {
		t.Fatalf("expected '123.45', got %v", result)
	}
}

func TestValidateAndFormatInput_EmptyInput(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("\n"))
	_, err := ValidateAndFormatInput()
	if err == nil || err.Error() != "input value can't be empty" {
		t.Fatalf("expected 'input value can't be empty' error, got %v", err)
	}
}

func TestValidateAndFormatInput_InvalidInput(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("abc\n"))
	_, err := ValidateAndFormatInput()
	if err == nil || err.Error() != "invalid input format" {
		t.Fatalf("expected 'invalid input format' error, got %v", err)
	}
}

func TestValidateAndFormatInput_ValidNegativeInput(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("-123.45\n"))
	result, err := ValidateAndFormatInput()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "-123.45" {
		t.Fatalf("expected '-123.45', got %v", result)
	}
}

func TestValidateAndFormatInput_ReadError(t *testing.T) {
	reader = bufio.NewReader(&errorReader{})
	_, err := ValidateAndFormatInput()
	if err == nil || err.Error() != "could not read input simulated read error" {
		t.Fatalf("expected 'could not read input simulated read error', got %v", err)
	}
}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}
