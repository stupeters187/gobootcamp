package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func main()  {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("program broke")
	}

	defer f.Close()

	//rdr := strings.NewReader("some string")

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("program broke")
	}

	fmt.Println(string(bs))
}