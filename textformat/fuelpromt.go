package textformat

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInput() (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read: %v", err)
	}

	return ValidateInput(input)
}

func ValidateInput(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", fmt.Errorf("entered value can't be empty")
	}

	punctuationRegEx := regexp.MustCompile(`^-?\d+([.,]?\d+)?$`)
	if !punctuationRegEx.MatchString(input) {
		return "", fmt.Errorf("entered value has invalid format")
	}

	return strings.ReplaceAll(input, ",", "."), nil
}

func PrintFuelConsumption() {
	fmt.Println("Enter fuel consumption per 100 kilometers.")
	drawPrompt()
}

func PrintFuelPrice() {
	fmt.Println("Enter fuel price per liter.")
	drawPrompt()
}

func drawPrompt() {
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")
}
