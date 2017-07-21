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
	err := gist.PostToGist(*description, *file)
	if err != nil {
		log.Fatal(err.Error())
	}
}
