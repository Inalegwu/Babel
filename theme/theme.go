package theme

import (
	"bufio"
	"encoding/json"
	"io/fs"

	"github.com/Inalegwu/babel/error"
)

type Theme struct {
	Name        string            `json:"name"`
	Theme_type  string            `json:"type"`
	Colors      map[string]string `json:"colors"`
	TokenColors []map[string]any  `json:"tokenColors"`
}

func New(scanner *bufio.Scanner, fileStat fs.FileInfo) Theme {
	var themeConfig Theme

	scanner.Split(bufio.ScanLines)

	var fileAsString []byte

	// for each line of input
	// This is so far the only way I can
	// achieve this , I don't know why
	for scanner.Scan() {
		// extrat the bytes
		line := scanner.Bytes()
		// for each substring in the line
		for i := 0; i < len(line); i++ {
			// append the item to the fileAsString
			// variable
			fileAsString = append(fileAsString, line[i])
		}
	}

	err := json.Unmarshal(fileAsString, &themeConfig)

	error.HandleError(err)

	return themeConfig
}
