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
	fuelConsumption = strings.Replace(fuelConsumption, "\n", "", -1)

	fmt.Println("Enter fuel price per liter.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")

	fuelPrice, _ = reader.ReadString('\n')
	fuelPrice = strings.Replace(fuelPrice, "\n", "", -1)

	literPrice := calculateKilometerPrice(fuelConsumption, fuelPrice)

	fmt.Printf("Kilometer price of fuel is: %.4f €", literPrice)
}

func calculateKilometerPrice(fuelConsumption, fuelPrice string) float64 {
	fuelConsumptionFloat, _ := strconv.ParseFloat(fuelConsumption, 64)
	fuelPriceFloat, _ := strconv.ParseFloat(fuelPrice, 64)

	var literPrice float64

	if fuelConsumptionFloat != 0 && fuelPriceFloat != 0 {
		literPrice = fuelConsumptionFloat * fuelPriceFloat / averageKilometers
	}

	return literPrice
}
