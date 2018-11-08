package main

import "fmt"

var x = map[int]string{
	1: "stu",
	2: "zobo",
}

func getIndex(index int) string {
	return x[index]
}

func main()  {

	delete(x, 1)
	fmt.Println(getIndex(2))
}
