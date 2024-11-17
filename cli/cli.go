package cli

import (
	"edchja.de/fuel-calc/logic"
	"edchja.de/fuel-calc/textformat"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func FuelCalcV2() {
	app := &cli.App{
		Name: "fuel-calc",
		Usage: "Calculate the average price of fuel per 100 kilometers.\n" +
			"The first argument represents the fuel consumption in liters per 100km.\n" +
			"The second argument represents the price or fuel in euros (€).\n\n" +
			"It is also possible to use the command without arguments.\n" +
			"In this case, the program will prompt for the fuel consumption and the fuel price.\n" +
			"Allowed separators are: [.,]",
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

			fmt.Printf(logic.CalculatedResponse, logic.CalculateConsumption(fuelConsumption, fuelPrice))

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
