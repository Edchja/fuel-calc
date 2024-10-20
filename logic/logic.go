package logic

import (
	"strconv"
)

const averageKilometers = 100

func CalculateKilometerPrice(fuelConsumption, fuelPrice string) (float64, error) {
	fuelConsumptionFloat, err := strconv.ParseFloat(fuelConsumption, 64)
	if err != nil {
		return 0, err
	}

	fuelPriceFloat, err := strconv.ParseFloat(fuelPrice, 64)
	if err != nil {
		return 0, err
	}

	var literPrice float64

	if fuelConsumptionFloat != 0 && fuelPriceFloat != 0 {
		literPrice = (fuelConsumptionFloat * fuelPriceFloat) / averageKilometers
	}

	return literPrice, nil
}
