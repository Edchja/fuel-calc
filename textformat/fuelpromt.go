package textformat

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ValidateAndFormatInput() (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read input %v", err)
	}
	input = strings.TrimSpace(input)

	if input == "" {
		return "", fmt.Errorf("input value can't be empty")
	}

	punctuationRegEx := regexp.MustCompile(`^-?\d+([.,]?\d+)?$`)
	if !punctuationRegEx.MatchString(input) {
		return "", fmt.Errorf("invalid input format")
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
