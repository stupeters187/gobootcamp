package main

import (
	"os"
	"log"
	"io/ioutil"
)

func main()  {
	dst, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln("program broke")
	}

	defer dst.Close()

	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("program broke")
	}

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("program broke")
	}

	str := []byte(bs)
	_, err = dst.Write(str)
	if err != nil {
		log.Fatalln("failure writing to file", err.Error())
	}
}
