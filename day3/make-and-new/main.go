package main

import "fmt"

func main() {
	var y int
	yPtr := &y
	*yPtr = y

	xPtr := new([100]byte)

	fmt.Println(yPtr)
	fmt.Println(y)
	fmt.Println(*yPtr)

	fmt.Println(xPtr)
}
