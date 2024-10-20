package main

import (
	"edchja.de/fuel-calc/textformat"
	"fmt"

	"edchja.de/fuel-calc/logic"
)

func main() {
	var fuelPrice string
	var fuelConsumption string

	textformat.PrintFuelConsumption()
	fuelConsumption, err := textformat.ValidateAndFormatInput(fuelConsumption)
	if err != nil {
		fmt.Printf("An error occoured: %s\n", err)
		return
	}

	textformat.PrintFuelPrice()
	fuelPrice, err = textformat.ValidateAndFormatInput(fuelPrice)
	if err != nil {
		fmt.Printf("An error occoured: %s\n", err)
		return
	}

	literPrice, err := logic.CalculateKilometerPrice(fuelConsumption, fuelPrice)
	if err != nil {
		fmt.Printf("An error occoured: %s\n", err)
		return
	}

	fmt.Printf("Kilometer price of fuel is: %.4f €", literPrice)
}
