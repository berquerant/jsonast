package jsonast_test

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/berquerant/jsonast"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const (
		inputFilename = "./testdata/input.json"
		wantFilename  = "./testdata/output.json"
	)

	inputFile, err := os.Open(inputFilename)
	if !assert.Nil(t, err) {
		return
	}
	defer inputFile.Close()
	wantFile, err := os.Open(wantFilename)
	if !assert.Nil(t, err) {
		return
	}
	defer wantFile.Close()

	wantBytes, err := io.ReadAll(wantFile)
	if !assert.Nil(t, err) {
		return
	}
	var want any
	if !assert.Nil(t, json.Unmarshal(wantBytes, &want)) {
		return
	}

	parsed, err := jsonast.Parse(inputFile)
	if !assert.Nil(t, err) {
		return
	}
	parsedBytes, err := json.Marshal(parsed)
	if !assert.Nil(t, err) {
		return
	}
	var got any
	if !assert.Nil(t, json.Unmarshal(parsedBytes, &got)) {
		return
	}

	assert.Equal(t, want, got)
}
