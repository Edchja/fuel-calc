package terminal

import (
	"edchja.de/fuel-calc/logic"
	"edchja.de/fuel-calc/textformat"
	"fmt"
)

func FuelCalcV1() {
	textformat.PrintFuelConsumption()
	fuelConsumption, err := textformat.ReadInput()
	if err != nil {
		logic.HandleError(err)
	}
	fuelConsumption, err = textformat.ValidateInput(fuelConsumption)
	if err != nil {
		logic.HandleError(err)
	}

	textformat.PrintFuelPrice()
	fuelPrice, err := textformat.ReadInput()
	if err != nil {
		logic.HandleError(err)
	}
	fuelPrice, err = textformat.ValidateInput(fuelPrice)
	if err != nil {
		logic.HandleError(err)
	}

	literPrice := logic.CalculateConsumption(fuelConsumption, fuelPrice)

	fmt.Printf(logic.CalculatedResponse, literPrice)
}
