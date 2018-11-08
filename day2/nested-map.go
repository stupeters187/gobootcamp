package main

import "fmt"

func main()  {

	x := map[string]map[string]string{
		"H": map[string]string{
			"name": "stu",
			"id": "187",
			"other": "no",
		},
		"B": map[string]string{
			"name": "zobo",
			"id": "69",
			"other": "yes",
		},
	}
	if el, ok := x["H"]; ok {
		fmt.Println(el["name"], el["id"])
	}
}
