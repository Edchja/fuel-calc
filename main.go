package main

import (
	"edchja.de/fuel-calc/logic"
	"edchja.de/fuel-calc/textformat"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const calculatedResponse = "The price per kilometer is: %.2f €"

func main() {
	// If no arguments are passed, v1 is used.
	if len(os.Args) == 1 {
		fuelCalcV1()
	}

	// If arguments are passed, v2 is used.
	if len(os.Args) > 1 {
		fuelCalcV2()
	}
}

func fuelCalcV2() {
	app := &cli.App{
		Name: "fuel-calc",
		Usage: "Calculate the price of fuel per kilometer.\n" +
			"The first argument is the fuel consumption in liters per 100km.\n" +
			"The second argument is the price of fuel in €.\n\n" +
			"It is also possible to use the command without arguments. " +
			"In this case, the program will ask for the fuel consumption and the fuel price." +
			"Allowed separators are: . ,",
		Version: "v2.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Edgar Timakin",
				Email: "edgartimakin+git@gmail.com",
			},
		},
		Copyright: "(c) 2024 Edgar Timakin",
		UsageText: "fuel-calc <fuel consumption> <fuel price>\nfuel-calc 16,9 1.99",
		Action: func(cCtx *cli.Context) error {
			var err error

			if cCtx.NArg() > 2 {
				return cli.Exit("too many arguments", -1)
			}

			// Example prompt: fuel-calc 17 1.99
			fuelConsumption := cCtx.Args().First()
			fuelPrice := cCtx.Args().Get(1)

			fuelConsumption, err = textformat.ValidateInput(fuelConsumption)
			if err != nil {
				return cli.Exit(err, -1)
			}

			fuelPrice, err = textformat.ValidateInput(fuelPrice)
			if err != nil {
				return cli.Exit(err, -1)
			}

			fmt.Printf(calculatedResponse, calculateConsumption(err, fuelConsumption, fuelPrice))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "Prints the current version of `fuel-calc`",
		Aliases: []string{"v"},
	}
}

func fuelCalcV1() {
	textformat.PrintFuelConsumption()
	fuelConsumption, err := textformat.ReadInput()
	if err != nil {
		handleError(err)
	}
	fuelConsumption, err = textformat.ValidateInput(fuelConsumption)
	if err != nil {
		handleError(err)
	}

	textformat.PrintFuelPrice()
	fuelPrice, err := textformat.ReadInput()
	if err != nil {
		handleError(err)
	}
	fuelPrice, err = textformat.ValidateInput(fuelPrice)
	if err != nil {
		handleError(err)
	}

	literPrice := calculateConsumption(err, fuelConsumption, fuelPrice)

	fmt.Printf(calculatedResponse, literPrice)
}

func calculateConsumption(err error, fuelConsumption string, fuelPrice string) float64 {
	fuelConsumptionFloat, err := logic.ParseToFloat(fuelConsumption)
	if err != nil {
		handleError(err)
	}

	fuelPriceFloat, err := logic.ParseToFloat(fuelPrice)
	if err != nil {
		handleError(err)
	}

	literPrice, err := logic.CalculateKilometerPrice(fuelConsumptionFloat, fuelPriceFloat)
	if err != nil {
		handleError(err)
	}
	return literPrice
}

func handleError(err error) {
	fmt.Printf("An error occoured: %s\n", err)
	os.Exit(1)
}
