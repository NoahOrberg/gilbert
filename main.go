package main

import (
	"flag"
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

	err := gist.PostToGist(*description, *file, isBasic)
	if err != nil {
		log.Fatal(err.Error())
	}
}
