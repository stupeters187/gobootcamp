package main

import "fmt"

func milesToKm(miles float64) float64 {
	return miles / 0.621369647819236
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * .5556
}

func lbToKg(lb float64) float64 {
	return lb / 2.2
}

func main()  {

	fmt.Println(milesToKm(1))
	fmt.Println(fahrenheitToCelsius(-20))
	fmt.Println(lbToKg(52.9109))
}
