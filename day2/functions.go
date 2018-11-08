package main

import "fmt"git

func minimum(xs []float64) float64 {

	min := xs[0]

	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}

func average(xs []float64) float64 {

	total := 0.0

	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main()  {

	xs := []float64{90, 30, 22, 87}
	n := minimum(xs)
	fmt.Println(n)
	fmt.Println(average(xs))
}
