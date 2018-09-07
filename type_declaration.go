package main

import "fmt"

type celsius int
type fahrenheit int
type kelvin celsius

const freeze celsius = 0
const boil celsius = 100

var temp celsius

func main() {
	var room int = 25

	// Can not be set because the types are not matching
	// temp = room

	// You can cast types if the underlying type is the same
	temp = celsius(room)

	fmt.Println("Temp:", temp)

	f := celsiusToFahrenheit(temp)

	fmt.Printf("Fahrenheit: %v, [%T]\n", f, f)

	var fZero fahrenheit = 0
	c := fahrenheitToCelsius(fZero)

	fmt.Printf("Celsius: %v, [%T]\n", c, c)

	var fMuch fahrenheit = 100
	var klein kelvin

	klein = kelvin(fMuch)
	fmt.Printf("Kelvin: %v, [%T]\n", klein, klein)

	// Run numeric assignments
	assignNumerics()
}

func celsiusToFahrenheit(c celsius) fahrenheit {
	f := c*9/5 + 32

	return fahrenheit(f)
}

func fahrenheitToCelsius(f fahrenheit) celsius {
	c := (f - 32) * 5 / 9

	return celsius(c)
}

func assignNumerics() {
	var pi float32 = 3.74
	var num int

	num = int(pi)

	fmt.Println(num)
}
