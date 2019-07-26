# go-browsercookie

- Port [Browsercookie](https://pypi.org/project/browsercookie/) from Python to Golang

## Install

```golang
go get github.com/onyas/go-browsercookie
```

## Usage

```

package main

import (
	"github.com/onyas/go-browsercookie"
	"log"
)

func main() {
	cookieJar, error := browsercookie.Chrome("https://google.com")
	if error != nil {
		log.Fatal(error)
	}

	log.Println(cookieJar)
}

```