package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
)

func main()  {
	f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln("program broke")
	}

	defer f.Close()

	rdr := strings.NewReader("some string")

	bs, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatalln("program broke")
	}

	fmt.Println(string(bs))
}