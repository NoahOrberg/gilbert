package gist

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/NoahOrberg/gilbert/config"
)

var ErrCouldNotLoad = errors.New("could not load gist")

type Gist struct {
	Files map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

type Payload struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type Response struct {
	HTMLURL string `json:"html_url"`
}

func createPayloadByContent(description string, g *Gist) Payload {
	payload := Payload{
		Description: description,
		Public:      false,
	}

	payload.Files = g.Files

	return payload
}

func PostToGistByContent(description, filename, content string) (string, error) {
	g := &Gist{
		Files: map[string]File{
			filename: File{
				Content: content,
			},
		},
	}

	res, err := PostToGist(description, g)
	if err != nil {
		return "", err
	}
	return res.HTMLURL, nil
}

func PostToGist(description string, g *Gist) (*Response, error) {
	conf := config.GetConfig()
	url := conf.GistURL

	p := createPayloadByContent(description, g)
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+conf.GistToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res *Response
	if resp.StatusCode == http.StatusCreated {
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}
	}

	return res, nil

}

func GetGist(id string) (*Gist, error) {
	var gist *Gist
	conf := config.GetConfig()
	url := conf.GistURL

	req, err := http.NewRequest(
		"GET",
		url+"/"+id,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+conf.GistToken)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, ErrCouldNotLoad
	}

	if err := json.NewDecoder(resp.Body).Decode(&gist); err != nil {
		return nil, err
	}

	return gist, nil
}

func DeleteGist(id string) error {
	conf := config.GetConfig()
	url := conf.GistURL

	req, err := http.NewRequest(
		"DELETE",
		url+"/"+id,
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+conf.GistToken)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func PatchGist(id string, gist Gist) (*Response, error) {
	conf := config.GetConfig()
	url := conf.GistURL

	payload, err := json.Marshal(gist)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PATCH",
		url+"/"+id,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+conf.GistToken)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var res *Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}
