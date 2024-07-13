package jsonast_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/berquerant/jsonast"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const dir = "./testdata"

	for i := range 2 {
		var (
			input  = fmt.Sprintf("%s/input_%d.json", dir, i)
			output = fmt.Sprintf("%s/output_%d.json", dir, i)
			title  = fmt.Sprintf("%s_%s", input, output)
		)
		t.Run(title, func(t *testing.T) {
			inputFile, err := os.Open(input)
			if !assert.Nil(t, err) {
				return
			}
			defer inputFile.Close()
			wantFile, err := os.Open(output)
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
		})
	}
}
