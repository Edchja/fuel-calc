package logic

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	averageKilometers  = 100
	CalculatedResponse = "The price per kilometer is: %.2f €"
)

func CalculateConsumption(fuelConsumption string, fuelPrice string) float64 {
	fuelConsumptionFloat, err := parseToFloat(fuelConsumption)
	if err != nil {
		HandleError(err)
	}

	fuelPriceFloat, err := parseToFloat(fuelPrice)
	if err != nil {
		HandleError(err)
	}

	literPrice := (fuelConsumptionFloat * fuelPriceFloat) / averageKilometers

	return literPrice
}

func parseToFloat(input string) (float64, error) {
	if input == "" {
		return 0, errors.New("entered value can't be empty")
	}

	inputFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("entered value invalid")
	}

	if inputFloat < 0 {
		return 0, errors.New("entered value can't be negative")
	}

	return inputFloat, nil
}

func HandleError(err error) {
	fmt.Printf("An error occoured: %s\n", err)
	os.Exit(1)
}
