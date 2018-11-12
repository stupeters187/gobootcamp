package main

import "fmt"

func zero(xptr *int)  {
	fmt.Println(xptr, *xptr)
	*xptr = 0
}

func main()  {
	x := 5
	zero(&x)
	fmt.Println(&x)
	fmt.Println(x) // x is 0
}