package main

import (
	"fmt"
	"strconv"
)

var measurement = map[int]string{
	1: "Kelvin",
	2: "Celsius",
	3: "Fahrenheit",
}

func convertKelvin(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := (((temperatureInput - 273.15) * 9) / 5) + 32
	resultCelsius := temperatureInput - 273.15
	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := (temperatureInput * 1.8) + 32
	resultKelvin := temperatureInput + 273.15
	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultKelvin := (temperatureInput-32)*5/9 + 273.15
	resultCelsius := (temperatureInput - 32) * 5 / 9
	//Do not remove this line
	return resultKelvin, resultCelsius
}

// Helper function to format output to 2 decimal place
func toTwoDecimal(input1 float64, input2 float64) (string, string) {
	res1 := fmt.Sprintf("%.2f", input1)
	res2 := fmt.Sprintf("%.2f", input2)
	return res1, res2
}

func main() {
	var temperatureChoiceInput, temperatureInput string
	var temperatureChoice int
	var temperature float64

	// Validate that user only input integer 1, 2 or 3
	for {
		fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
		fmt.Scanln(&temperatureChoiceInput)
		v, err := strconv.Atoi(temperatureChoiceInput)
		if err != nil {
			fmt.Println("Invalid choice.")
		} else {
			name, ok := measurement[v]
			if ok {
				temperatureChoice = v
				fmt.Println("You have selected: ", name)
				break
			} else {
				fmt.Println("Invalid choice.")
			}
		}
	}

	// Validate that user only input float64
	for {
		fmt.Println("Enter the temperature: ")
		fmt.Scanln(&temperatureInput)
		v, err := strconv.ParseFloat(temperatureInput, 64)
		if err != nil {
			fmt.Println("Please enter a valid temperature.")
		} else {
			temperature = v
			break
		}
	}

	if temperatureChoice == 1 {
		//Insert Code here
		resultFahrenheit, resultCelsius := toTwoDecimal(convertKelvin(temperature))
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Celsius: ", resultCelsius)
	} else if temperatureChoice == 2 {
		//Insert Code here
		resultFahrenheit, resultKelvin := toTwoDecimal(convertCelsius(temperature))
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Kelvin: ", resultKelvin)
	} else if temperatureChoice == 3 {
		//Insert Code here
		resultKelvin, resultCelsius := toTwoDecimal(convertFahrenheit(temperature))
		//DO not remove this
		fmt.Println("Kelvin: ", resultKelvin, " Celsius: ", resultCelsius)
	} else {
		fmt.Println("No Conversion")
	}

}
