package main

import (
	"io"
	"bufio"
	"os"
	"log"
	"fmt"
)

func LongestWord(rdr io.Reader) string {
	currentLongestWord := ""
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > len(currentLongestWord) {
			currentLongestWord = word
		}
	}
	return currentLongestWord
}

func main() {
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	longest := LongestWord(src)
	fmt.Println(longest)

}