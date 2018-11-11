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

	fmt.Fprintln(os.Stderr, "Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Fprintln(os.Stderr, "Please enter your email: ")
	var email string
	fmt.Scanln(&email)
	fmt.Fprintln(os.Stderr, "Please enter your website: ")
	var website string
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