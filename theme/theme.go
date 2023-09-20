package theme

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"log"

	"github.com/Inalegwu/vsconvert/error"
)

type Theme struct {
	Name        string              `json:"name"`
	Theme_type  string              `json:"type"`
	Colors      map[string]string   `json:"colors"`
	TokenColors []map[string]string `json:"tokenColors"`
}

func New(scanner *bufio.Scanner, fileStat fs.FileInfo) Theme {
	log.Printf("Reading %s", fileStat.Name())
	var themeConfig Theme

	scanner.Split(bufio.ScanLines)

	fileAsString, ff := json.Marshal(scanner.Bytes())

	error.HandleError(ff)

	err := json.Unmarshal(fileAsString, &themeConfig)

	error.HandleError(err)

	return themeConfig
}
