package main

import (
	"fmt"
)

func main()  {

	xsx := []int{
		12, 35, 78, 90,
		55, 98, 69, 41,
		34, 10, 55, 99, 1,
	}

	min := xsx[0]

	for _, v := range xsx {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
}
