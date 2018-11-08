package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {

	from := os.Args[1]
	to := os.Args[2]

	switch {
	case strings.HasSuffix(from, "mi"):
		fmt.Println("FROM MILES")
	}

	fmt.Println(from, to)
}