package main

import "fmt"

func swapAndRotate(args ...*int)  {

	firstArg := *args[0]
	for i := 0; i < len(args)-1 ; i++ {
		*args[i] = *args[i+1]
	}
	*args[len(args)-1] = firstArg

}

func main()  {
	x := 1
	y := 2
	z := 3
	swapAndRotate(&x, &y, &z)
	fmt.Println(x,y,z)
}
