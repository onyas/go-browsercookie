package main

import (
	".."
	"fmt"
	"log"
)

func main() {
	cookieJar, err := browsercookie.LoadCookieJarFromChrome("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cookieJar)
}
