package main

import (
	"fmt"
	"strings"
	"crypto/md5"
	"io"
	"encoding/hex"
	"os"
	"log"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}

func main()  {

	var name string
	var email string
	var website string
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Please enter your email: ")
	fmt.Scanln(&email)
	fmt.Println("Please enter your website: ")
	fmt.Scanln(&website)

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	gravatarHash := getGravatarHash(email)
	fmt.Fprintf(file,`<!DOCTYPE html>
<html>
	<head>
	</head>
	<body>
		<h1>Welcome to `+ name +`'s Profile</h1>
		<img src="http://www.gravatar.com/avatar/` + gravatarHash + `?d=identicon">
		<li>
			<ul><a href="mailto:`+ email +`">email ` + name + `</a></ul>
			<ul><a href="http://` + website + `">` + name + `'s website</a></ul>
		</li>
	</body>
</html>
`)
}