package main

import (

	"strings"
	"os"
	"log"
	"io/ioutil"
)

func main() {

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
	str := string([]byte(bs))
	bs = []byte(strings.ToUpper(str))
	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("failure writing to file", err.Error())
	}
}
