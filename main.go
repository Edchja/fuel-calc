package main

import (
	"fmt"

	"edchja.de/fuel-calc/logic"
)

func main() {
	var fuelPrice string
	var fuelConsumption string

	logic.PrintFuelConsumption()
	fuelConsumption, err := logic.FormatString(fuelConsumption)
	if err != nil {
		fmt.Printf("An error occoured: %s", err)
	}

	logic.PrintFuelPrice()
	fuelPrice, err = logic.FormatString(fuelPrice)
	if err != nil {
		fmt.Printf("An error occoured: %s", err)
	}

	literPrice, err := logic.CalculateKilometerPrice(fuelConsumption, fuelPrice)
	if err != nil {
		fmt.Printf("An error occoured: %s", err)
	}

	fmt.Printf("Kilometer price of fuel is: %.4f €", literPrice)
}
