package main

import (
	"strings"
	"fmt"
)

func Wordcount(str string) map[string]int {
	strMap := make(map[string]int)
	words := strings.Fields(str)
	for _, v := range words {
		strMap[v]++
	}
	return strMap
}


func main()  {

	str := "test test test"
	result := Wordcount(str)
	fmt.Println(result)
}
