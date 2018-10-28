package main

import ( "fmt"
		 "strings"
		 "strconv"
		 )

var str = "string"

func main()  {

	str = "newstring"
	fmt.Println(str)
	fmt.Println(str[0])
	fmt.Println(strings.Replace("Hello <NAME>", "<NAME>", "stu", 1))
	fmt.Println(strconv.ParseInt(
		"AF0875", // string to parse
		16, // base to convert to
		64)) //bitsize
	fmt.Println(strconv.FormatInt(123, 16))
}
