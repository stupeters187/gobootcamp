package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"fmt"
)

func Wordcount(str string) map[string]int {
	strMap := map[string]int{}
	for _, word := range strings.Fields(str) {
		strMap[word]++
	}
	return strMap
}

func main() {

	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("program broke")
	}

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("program broke")
	}
	str := string([]byte(bs))
	result := Wordcount(str)
	fmt.Println(result)
}
