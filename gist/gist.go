package gist

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/NoahOrberg/gilbert/config"
)

type Payload struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	File        map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

func createPayload(description, file string) (Payload, error) {
	payload := Payload{
		Description: description,
		Public:      false,
	}

	f, err := os.Open(file)
	if err != nil {
		errors.New("No such file or directory :" + file)
	}
	defer f.Close()

	var content string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content += scanner.Text()
		content += "\n"
	}

	filename := strings.Split(file, "/")[len(strings.Split(file, "/"))-1]

	payload.File = map[string]File{
		filename: File{Content: content},
	}

	return payload, nil
}

func PostToGist(description, file string) error {
	url := "https://api.github.com/gists"

	// create payload
	p, err := createPayload(description, file)
	if err != nil {
		return err
	}

	config := config.GetConfig()

	payload, err := json.Marshal(p)
	if err != nil {
		return err
	}

	log.Print(string(payload))

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+config.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Print(resp.Status)

	return nil
}
