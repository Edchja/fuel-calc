package textformat

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ValidateAndFormatInput(input string) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read input")
	}
	input = strings.TrimSpace(input)

	reg := regexp.MustCompile(`^-?\d+([.,]?\d+)?$`)
	if !reg.MatchString(input) {
		return "", fmt.Errorf("invalid input format")
	}

	return strings.NewReplacer(",", ".").Replace(input), nil
}

func PrintFuelConsumption() {
	fmt.Println("Enter fuel consumption per 100 kilometers.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")
}

func PrintFuelPrice() {
	fmt.Println("Enter fuel price per liter.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")
}
