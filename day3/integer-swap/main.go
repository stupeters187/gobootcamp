package main

import "fmt"

func swap(x, y *int)  {
	xPtr := *x
	yPtr := *y
	*x = yPtr
	*y = xPtr
}

func main()  {
	x := 1
	y := 2
	fmt.Printf("value of x: %d\n", x)
	fmt.Printf("value of y: %d\n", y)
	fmt.Println("running swap funtion...")
	swap(&x, &y)
	fmt.Printf("value of x: %d\n" ,x)
	fmt.Printf("value of y: %d\n" ,y)
}
