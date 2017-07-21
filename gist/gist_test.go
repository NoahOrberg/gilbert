package gist

import (
	"bufio"
	"os"
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

	actual, err := createPayload(description, file)
	require.NoError(err)

	assert.Equal(expected, actual)

}
