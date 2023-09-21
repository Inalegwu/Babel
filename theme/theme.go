package theme

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"log"

	"github.com/Inalegwu/babel/error"
)

type Theme struct {
	Name       string            `json:"name"`
	Theme_type string            `json:"type"`
	Colors     map[string]string `json:"colors"`
	// TokenColors map[string][]TokenConfig `json:"tokenColors"`
}

func New(scanner *bufio.Scanner, fileStat fs.FileInfo) Theme {
	log.Printf("Reading %s", fileStat.Name())
	var themeConfig Theme

	scanner.Split(bufio.ScanLines)

	var fileAsString []byte

	for scanner.Scan() {
		line := scanner.Bytes()
		for i := 0; i < len(line); i++ {
			fileAsString = append(fileAsString, line[i])
		}
	}

	err := json.Unmarshal(fileAsString, &themeConfig)

	error.HandleError(err)

	return themeConfig
}
