package main

import "fmt"

func increasingNumberGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main()  {

	numbers := increasingNumberGenerator()

	fmt.Println(numbers())
	fmt.Println(numbers())
}