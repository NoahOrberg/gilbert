package gist

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/NoahOrberg/gilbert/config"
)

func createPayload(description string, files []string) (string, error) {
	payload := `{"description":"` + description + `", "public":false, "files":{`
	for _, f := range files {
		splitFName := strings.Split(f, "/")
		fname := splitFName[len(splitFName)-1]
		file, err := os.Open(f)
		if err != nil {
			log.Print("No such file or directory :" + f)
			continue
		}
		defer file.Close()
		var content string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			content += scanner.Text()
			content += "\\n"
		}

		payload += `"` + fname + `":{"content":"` + content + `"}, `
	}
	payload = strings.Trim(payload, ", ")
	payload += "}}"

	return payload, nil
}

func PostToGist(description string, files []string) error {
	url := "https://api.github.com/gists"

	// create payload
	payload, err := createPayload(description, files)
	if err != nil {
		return err
	}

	// log
	log.Print(payload)

	config := config.GetConfig()

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(payload)),
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
