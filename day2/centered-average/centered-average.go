package main

import "fmt"

//Drop the highest number in the slice
func dropHighest(arr []float64) []float64 {
	min := arr[0]
	var index int
	for i, v := range arr {
		if min < v {
			min = v
			index = i
		}
	}
	arr[index] = arr[len(arr)-1]
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	return arr
}

func dropLowest(arr []float64) []float64 {
	min := arr[0]
	var index int
	for i, v := range arr {
		if min > v {
			min = v
			index = i
		}
	}
	arr[index] = arr[len(arr)-1]
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	return arr
}

func average(arr []float64) float64 {
	arr = dropLowest(arr)
	arr = dropHighest(arr)
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}

func main()  {
	arr := []float64{90,21,3,45,67,73,11}
	fmt.Println(average(arr))
}
