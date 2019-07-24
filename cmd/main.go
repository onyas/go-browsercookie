package main

import (
	".."
	"fmt"
	"log"
)

func main() {
	cookieJar, err := browsercookie.LoadCookieJarFromChrome("https://bing.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cookieJar)
}
