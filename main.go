package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const averageKilometers = 100

func main() {
	var fuelPrice string
	var fuelConsumption string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter fuel consumption per 100 kilometers.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")

	fuelConsumption, _ = reader.ReadString('\n')
	// convert CRLF to LF
	fuelConsumption = strings.Replace(fuelConsumption, "\r\n", "", -1)

	fmt.Println("Enter fuel price per liter.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")

	fuelPrice, _ = reader.ReadString('\n')
	fuelPrice = strings.Replace(fuelPrice, "\r\n", "", -1)

	literPrice, err := calculateKilometerPrice(fuelConsumption, fuelPrice)
	if err != nil {
		fmt.Printf("An error occoured: %s", err)
	}

	fmt.Printf("Kilometer price of fuel is: %.4f €", literPrice)
}

func calculateKilometerPrice(fuelConsumption, fuelPrice string) (float64, error) {
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
		literPrice = fuelConsumptionFloat * fuelPriceFloat / averageKilometers
	}

	return literPrice, nil
}
