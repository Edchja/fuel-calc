package logic

import (
	"errors"
	"strconv"
)

const averageKilometers = 100

func CalculateKilometerPrice(fuelConsumption, fuelPrice string) (float64, error) {
	if fuelConsumption == "" || fuelPrice == "" {
		return 0, errors.New("input values can't be empty")
	}

	fuelConsumptionFloat, err := strconv.ParseFloat(fuelConsumption, 64)
	if err != nil {
		return 0, errors.New("invalid fuel consumption value")
	}

	fuelPriceFloat, err := strconv.ParseFloat(fuelPrice, 64)
	if err != nil {
		return 0, errors.New("invalid fuel price value")
	}

	literPrice := (fuelConsumptionFloat * fuelPriceFloat) / averageKilometers

	return literPrice, nil
}
