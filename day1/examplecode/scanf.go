package main

import "fmt"

func main() {

	var s string
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s", &s)
	fmt.Println("Hello", s)
}

