package main

import (
	"io"
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
)

func LongestWord(rdr io.Reader) string {
	currentLongestWord := ""
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {

		wordOrManyWords := scanner.Text()
		wordOrManyWords = strings.Replace(wordOrManyWords,"?", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords,"--", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords,"-", " ", -1)
		wordOrManyWords = strings.Replace(wordOrManyWords,"/", " ", -1)

		for _, word := range strings.Fields(wordOrManyWords) {
			if len(word) > len(currentLongestWord) {
				currentLongestWord = word
			}
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