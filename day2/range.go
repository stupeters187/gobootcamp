package main

import "fmt"

var arr = []int{1,2,3,4}

func main()  {

	var num string
	fmt.Println("do you need the index? (y or n)")
	fmt.Scanf("%s", &num )

	if num == "y" {
		for i ,c := range arr {
			fmt.Println(i, c)
		}
	} else {
		for _ ,c := range arr {
			fmt.Println(c)
		}
	}

}