package main

import "fmt"

func fillWithOne(x *[100]byte)  {
	fmt.Printf("address in fillWithOne: %p\n", &x)
	for i := range x {
		x[i] = 1
	}
}

func main()  {
	x := [100]byte{}
	fmt.Printf("address in main: %p\n", &x)
	fillWithOne(&x)
	fmt.Println(x)
}
