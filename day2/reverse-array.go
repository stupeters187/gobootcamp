package main

import "fmt"

func reverse(arr []int) ([]int) {
	var new []int
	for i := len(arr)-1; i >= 0; i--{
		new = append(new, arr[i])
	}
	return new
}

func main()  {

	arr := []int{1,2,3,4,5,6,9,10,11,12}
	var newArr []int = reverse(arr)
	fmt.Println(newArr)

}
