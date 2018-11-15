package main

import "fmt"

type MyClass struct {
	x int
}

func myNewClass(x int) *MyClass {
	return &MyClass{
		x: x,
	}
}

func (this *MyClass) whatever() {
	fmt.Println(this.x)
}

func main()  {
	obj := myNewClass(5)
	obj.whatever()
}