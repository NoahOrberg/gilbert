package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/NoahOrberg/gilbert/gist"
)

func readFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var content string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content += scanner.Text()
		content += "\n"
	}

	return content, nil
}

func main() {
	var description = flag.String("d", "", "description")
	var file = flag.String("f", "", "upload file")
	flag.Parse()

	content, err := readFile(*file)
	if err != nil {
		log.Fatal(err.Error())
	}

	url, err := gist.PostToGistByContent(*description, *file, content)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(url)
}
