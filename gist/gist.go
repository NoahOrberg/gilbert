package gist

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

func createPayload(description, file string) (Payload, error) {
	payload := Payload{
		Description: description,
		Public:      false,
	}

	f, err := os.Open(file)
	if err != nil {
		return Payload{}, errors.New("No such file or directory :" + file)
	}
	defer f.Close()

	var content string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content += scanner.Text()
		content += "\n"
	}

	tempFilename := strings.Split(file, "/")
	filename := tempFilename[len(tempFilename)-1]

	payload.File = map[string]File{
		filename: File{Content: content},
	}

	return payload, nil
}

func PostToGist(description, file string, isBasic bool) error {
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

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}

	if isBasic {
		var username string
		fmt.Println("Please login")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		req.SetBasicAuth(username, string(password))
		fmt.Println("")
	} else {
		req.Header.Set("Authorization", "token "+config.Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Print(resp.Status)

	if resp.StatusCode == http.StatusCreated {
		var res Response
		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Print(res.HTMLURL)
	}

	if resp.StatusCode == http.StatusUnauthorized {
	}

	return nil
}
