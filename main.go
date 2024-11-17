package main

import (
	"edchja.de/fuel-calc/cli"
	"edchja.de/fuel-calc/terminal"
	"os"
)

func main() {
	// If no arguments are passed, v1 is used.
	if len(os.Args) == 1 {
		terminal.FuelCalcV1()
	}

	// If arguments are passed, v2 is used.
	if len(os.Args) > 1 {
		cli.FuelCalcV2()
	}
}
