package main

import (
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


func main()  {

	str := "test test test"
	result := Wordcount(str)
	fmt.Println(result)
}
