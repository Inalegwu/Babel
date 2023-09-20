package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Inalegwu/vsconvert/convert"
	"github.com/Inalegwu/vsconvert/error"
	"github.com/Inalegwu/vsconvert/theme"
)

func main() {
	log.SetPrefix("VSCONVERT::")
	filePath := os.Getenv("FILE")

	if filePath == "" {
		log.Fatal("Please provide a file path by passing FILE=`file_path` ")
	}

	// try opening file
	file, err := os.Open(filePath)

	error.HandleError(err)

	// initialize the converter with a ConverterType flag (int)
	converter := convert.New(1)

	// get file info to use for unmarshalling
	fileStat, err := file.Stat()
	error.HandleError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	theme := theme.New(scanner, fileStat)

	converter.Convert(theme)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
