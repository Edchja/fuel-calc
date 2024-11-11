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
	result, err := ReadInput()
	if err != nil {
		t.Fatalf("expected no error, got '%v'", err)
	}
	if result != "123.45" {
		t.Fatalf("expected '123.45', got '%v'", result)
	}
}

func TestValidateAndFormatInput_ValidInputWithComma(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("123,45\n"))
	result, err := ReadInput()
	if err != nil {
		t.Fatalf("expected no error, got '%v'", err)
	}
	if result != "123.45" {
		t.Fatalf("expected '123.45', got '%v'", result)
	}
}

func TestValidateAndFormatInput_EmptyInput(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("\n"))
	_, err := ReadInput()
	if err == nil || err.Error() != "entered value can't be empty" {
		t.Fatalf("expected 'entered value can't be empty' error, got '%v'", err)
	}
}

func TestValidateAndFormatInput_InvalidInput(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("abc\n"))
	_, err := ReadInput()
	if err == nil || err.Error() != "entered value has invalid format" {
		t.Fatalf("expected 'entered value has invalid format' error, got '%v'", err)
	}
}

func TestValidateAndFormatInput_ValidNegativeInput(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("-123.45\n"))
	result, err := ReadInput()
	if err != nil {
		t.Fatalf("expected no error, got '%v'", err)
	}
	if result != "-123.45" {
		t.Fatalf("expected '-123.45', got '%v'", result)
	}
}

func TestValidateAndFormatInput_ReadError(t *testing.T) {
	reader = bufio.NewReader(&errorReader{})
	_, err := ReadInput()
	if err == nil || err.Error() != "could not read: simulated read error" {
		t.Fatalf("expected 'could not read: simulated read error', got '%v'", err)
	}
}

func (e *errorReader) Read([]byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}
