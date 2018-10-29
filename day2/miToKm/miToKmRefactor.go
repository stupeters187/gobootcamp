package main

import (
	"fmt"
	"os"
	"strconv"
)

var miToKm float64 = 1.60934

func main()  {

	// get the number of miles from the user
	number, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")

	fmt.Printf("Miles: %f<br>\n", number)
	fmt.Printf("Kilometers: %f<br>\n", number * miToKm)

	fmt.Println("</body>")
	fmt.Println("</html>")
}
