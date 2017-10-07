package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/NoahOrberg/gilbert/gist"
)

func main() {
	var description = flag.String("d", "", "description")
	var file = flag.String("f", "", "upload file")
	flag.Parse()

	var isBasic bool
	if ok := gist.IsNotUndefinedToken(); !ok {
		isBasic = true
	}

	url, err := gist.PostToGistByFile(*description, *file, isBasic)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(url)
}
