package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Println(file)
	fmt.Fprintf(file, "Hello Readers of golangcode.com")
}