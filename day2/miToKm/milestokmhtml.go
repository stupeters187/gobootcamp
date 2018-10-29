package main

import "fmt"

var miToKm float64 = 1.60934

func main()  {

	// get the number of miles from the user
	fmt.Println("Please enter the number of miles to convert:")
	var number float64
	fmt.Scanln(&number)

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")

	fmt.Printf("Miles: %f<br>\n", number)
	fmt.Printf("Kilometers: %f<br>\n", number * miToKm)

	fmt.Println("</body>")
	fmt.Println("</html>")


}
