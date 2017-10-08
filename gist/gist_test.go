package gist

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePayload(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	file := "../testcase/test.txt"
	description := "DESC"

	f, err := os.Open(file)
	require.NoError(err)
	defer f.Close()

	var content string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content += scanner.Text()
		content += "\n"
	}

	expected := Payload{
		Description: "DESC",
		Public:      false,
		File: map[string]File{
			"test.txt": File{
				Content: content,
			},
		},
	}

	actual, err := createPayloadByFile(description, file)
	require.NoError(err)

	assert.Equal(expected, actual)

}

func TestGetGist(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	content := "AAA"
	filename := "BBB"

	url, err := PostToGistByContent("", filename, content)
	require.NoError(err)

	splittedURL := strings.Split(url, "/")
	id := splittedURL[len(splittedURL)-1]

	g, err := GetGist(id)
	require.NoError(err)

	c, ok := g.Files[filename]
	assert.True(ok)
	assert.Equal(content, c.Content)

	err = DeleteGist(id)
	require.NoError(err)
}

func TestPatchGist(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	content := "AAA"
	filename := "BBB"

	url, err := PostToGistByContent("", filename, content)
	require.NoError(err)

	splittedURL := strings.Split(url, "/")
	id := splittedURL[len(splittedURL)-1]

	newContent := content + "AA"
	ng := Gist{
		Files: map[string]File{
			filename: File{
				Content: newContent,
			},
		},
	}

	err = PatchGist(id, ng)
	require.NoError(err)

	g, err := GetGist(id)
	require.NoError(err)

	c, ok := g.Files[filename]
	assert.True(ok)
	assert.Equal(newContent, c.Content)

	err = DeleteGist(id)
	require.NoError(err)
}
