package main

import "fmt"

type MyClass struct {
	x int
}

func NewMyClass(x int) *MyClass {
	return &MyClass{
		x: x,
	}
}

func (this *MyClass) whatever() {
	fmt.Println(this.x)
}

func main()  {
	obj := NewMyClass(5)
	obj.whatever()
}