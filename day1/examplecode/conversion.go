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

	var container float64
	var choice int64
	fmt.Println("What would you like to convert?")
	fmt.Print("1 for distance, 2 for temp, 3 for weight: ")
	fmt.Scanf("%v", &choice)
	fmt.Print("Convert for which value?: ")
	fmt.Scanf("%v", &container)
	switch choice {
	case 1:
		fmt.Println(milesToKm(container))
	case 2:
		fmt.Println(fahrenheitToCelsius(container))
	case 3:
		fmt.Println(lbToKg(container))
	}
}
