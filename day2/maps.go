package main

import "fmt"

var x = make(map[string]int)

func getIndex(index string) int {
	return x[index]
}

func main()  {

	x["stu"] = 187
	x["zoe"] = 0

	v, ok := x["zoe"]

	fmt.Println(x["stu"])
	fmt.Println(v, ok)
	fmt.Println(getIndex("stu"))
}
