package logic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const averageKilometers = 100

var reader = bufio.NewReader(os.Stdin)

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
		literPrice = fuelConsumptionFloat * fuelPriceFloat / averageKilometers
	}

	return literPrice, nil
}

func FormatString(input string) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Replace(input, "\r\n", "", -1), nil
}

func PrintFuelConsumption() {
	fmt.Println("Enter fuel consumption per 100 kilometers.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")
}

func PrintFuelPrice() {
	fmt.Println("Enter fuel price per liter.")
	fmt.Println("------------------------------------------")
	fmt.Print("-> ")
}
