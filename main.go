package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {

	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}
	originUnit = strings.ToUpper(os.Args[1])
	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err := fmt.Scanln(&originValue)
		if err != nil {
			printError(errReadingInput)
		}
		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}
		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		_, err = fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			printError(errReadingInput)
		}

		/*
			Currently, we can't always be sure that the value assigned to shouldConvertAgain
			is going to be in the casing we expect. Let's fix this by making a small change to
			the condition on the last if statement inside the for loop. The condition
			should be a combination of calling the strings.ToUpper() function,
			passing it the result of calling the strings.TrimSpace()
			function with shouldConvertAgain as its argument.
			If the result of all of that is NOT equal to "Y", then
			the existing if block containing fmt.Println("Good bye!")
			followed by break should be run.*/

		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
