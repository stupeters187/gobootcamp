package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	//"sort"
)

func main() {

	//var longest int
	strMap := map[int][]string{}

	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("program broke")
	}

	defer src.Close()

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("program broke")
	}
	str := string(bs)

	for _, word := range strings.Fields(str) {
		fmt.Println(len(word))
		//length := len(word)
		strMap = append(strMap[length], word)
		//////fmt.Println(strMap)
		////var keys []int
		////for k := range strMap {
		////	keys = append(keys, k)
		////}
		////
		////sort.Ints(keys)
		////
		////for _, k := range keys {
		////	fmt.Println("Key:", k, "Value:", strMap[k])
		////}
		//fmt.Println(strMap)
	}
	//fmt.Println(longestWords)
}