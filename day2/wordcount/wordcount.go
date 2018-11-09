package main

import (
	"strings"
	"fmt"
)

func main()  {

	strMap := make(map[string]int)
	str := strings.Split("test test test", " ")
	for _, v := range str {
		strMap[v]++
	}
	fmt.Println(strMap["test"])
}
