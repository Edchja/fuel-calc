package logic

import (
	"errors"
	"strconv"
)

const averageKilometers = 100

func CalculateKilometerPrice(fuelConsumption, fuelPrice float64) (float64, error) {
	literPrice := (fuelConsumption * fuelPrice) / averageKilometers

	return literPrice, nil
}

func ParseToFloat(input string) (float64, error) {
	if input == "" {
		return 0, errors.New("input value can't be empty")
	}

	inputFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("invalid input value")
	}

	if inputFloat < 0 {
		return 0, errors.New("input value can't be negative")
	}

	return inputFloat, nil
}
