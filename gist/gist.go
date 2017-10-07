package gist

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/NoahOrberg/gilbert/config"
	"golang.org/x/crypto/ssh/terminal"
)

type Payload struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	File        map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

type Response struct {
	HTMLURL string `json:"html_url"`
}

func IsNotUndefinedToken() bool {
	config := config.GetConfig()
	if config.Token == "" {
		return false
	} else {
		return true
	}
}

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

func createPayloadByFile(description, file string) (Payload, error) {
	payload := Payload{
		Description: description,
		Public:      false,
	}

	content, err := readFile(file)
	if err != nil {
		return Payload{}, err
	}

	tempFilename := strings.Split(file, "/")
	filename := tempFilename[len(tempFilename)-1]

	payload.File = map[string]File{
		filename: File{Content: content},
	}

	return payload, nil
}

func createPayloadByContent(description, filename, content string) (Payload, error) {
	payload := Payload{
		Description: description,
		Public:      false,
	}

	payload.File = map[string]File{
		filename: File{Content: content},
	}

	return payload, nil
}

func PostToGistByContent(description, filename, content string) (string, error) {
	url := "https://api.github.com/gists"

	// create payload
	p, err := createPayloadByContent(description, filename, content)
	if err != nil {
		return "", err
	}

	config := config.GetConfig()

	payload, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "token "+config.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res Response
	if resp.StatusCode == http.StatusCreated {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res.HTMLURL, nil

}

func PostToGistByFile(description, file string, isBasic bool) (string, error) {
	url := "https://api.github.com/gists"

	// create payload
	p, err := createPayloadByFile(description, file)
	if err != nil {
		return "", err
	}

	config := config.GetConfig()

	payload, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return "", err
	}

	if isBasic {
		var username string
		fmt.Println("Please login")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		req.SetBasicAuth(username, string(password))
		fmt.Println("")
	} else {
		req.Header.Set("Authorization", "token "+config.Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res Response
	if resp.StatusCode == http.StatusCreated {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res.HTMLURL, nil
}
