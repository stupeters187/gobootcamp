package main

import "fmt"

func main()  {

	outer:
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("if statemant")
		}
		fmt.Println("loop", i)

		switch {
		case i == 5:
			fmt.Println("will print")
			break outer

		case i == 9:
			fmt.Println("will print???")
		}
	}
}
