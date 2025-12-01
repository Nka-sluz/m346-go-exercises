package main

import "fmt"

func convertCelsiusToFahrenheit(c float64) float64 {
	return c*(9.0/5.0) + 32
}
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * (5.0 / 9.0)
}

type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*(9.0/5.0) + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * (5.0 / 9.0))
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(20)) // 68
	fmt.Println(convertCelsiusToFahrenheit(25)) // 77
	fmt.Println(convertCelsiusToFahrenheit(30)) // 86
	fmt.Println(convertFahrenheitToCelsius(68)) // 20
	fmt.Println(convertFahrenheitToCelsius(77)) // 25
	fmt.Println(convertFahrenheitToCelsius(86)) // 30

	var cozy Celsius = 23.0
	fmt.Println(cozy.ConvertToFahrenheit()) // 73.4

	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius()) // -9.277777777777779

	// übersichtlicher:
	var c Celsius = 23
	fmt.Println(c.ConvertToFahrenheit().ConvertToCelsius()) // 23.000000000000004
}
