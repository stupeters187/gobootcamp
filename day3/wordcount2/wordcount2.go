package main

import (
	"strings"
	"fmt"
)

func Wordcount(str string) []string {
	strArr := strings.Fields(str)
	return strArr
}

func main()  {

	str := "this is a string to be returned"
	fmt.Println(Wordcount(str))
}
