package main

import "fmt"

func main()  {
	var x = []int{1,2,3}

	var y = make([]int, 3)
	var z = make([]int, 10, 15)

	yz := make([]int, 4)
	copy(yz[1:], x)

	ys := make([]int, 4)
	ys[0] = 8

	x = append(x, ys[0:]...)

	fmt.Println(x, y, z, yz)
	fmt.Println(x[1:])
	fmt.Println(x[:2])
}