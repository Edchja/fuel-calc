package main

import (
	"edchja.de/fuel-calc/textformat"
	"fmt"
	"os"

	"edchja.de/fuel-calc/logic"
)

func main() {
	textformat.PrintFuelConsumption()
	fuelConsumption, err := textformat.ValidateAndFormatInput()
	if err != nil {
		handleError(err)
	}

	textformat.PrintFuelPrice()
	fuelPrice, err := textformat.ValidateAndFormatInput()
	if err != nil {
		handleError(err)
	}

	literPrice, err := logic.CalculateKilometerPrice(fuelConsumption, fuelPrice)
	if err != nil {
		handleError(err)
	}

	fmt.Printf("Kilometer price of fuel is: %.2f €", literPrice)
}

func handleError(err error) {
	fmt.Printf("An error occoured: %s\n", err)
	os.Exit(1)
}
