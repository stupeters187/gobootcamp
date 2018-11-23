package main

import (
	"os"
	"log"
)

func main()  {
	f, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln("program broke")
	}

	defer f.Close()

	str := "txt"
	bs := []byte(str)

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("failure writing to file", err.Error())
	}
}