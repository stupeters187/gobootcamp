package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strings.ToUpper(string(line[0])) + string(line[1:]))
	}
}