package main

import (
	"io"
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
)

func Wordcount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		counts[word]++
	}
	return counts
}

func main() {
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	counts := Wordcount(src)
	bigCount := counts["whale"] + counts["whale,"] + counts["whale."]
	fmt.Println("Number of whales:", bigCount)
}