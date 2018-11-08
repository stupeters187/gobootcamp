package main

import "fmt"

func main()  {

	var code string
	fmt.Scanln(&code)

	states := map[string]string{
		"CA": "california",
	}

	fmt.Println(states[code])
}
