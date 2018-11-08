package main

import "fmt"

func main()  {

	var x = [5]float64{79, 90, 43, 92, 7}

	var total float64

	for _, v := range x{
		total += v
	}
	fmt.Println(total / float64(len(x)))
}

