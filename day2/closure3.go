package main

import "fmt"

func visit(numbers []int, callback func(int))  {
	for _, n := range numbers{
		callback(n)
	}
}

func main() {
	x := 0
	visit([]int{1,2,3,4}, func(i int) {
		fmt.Println(x + i)
	})
}
