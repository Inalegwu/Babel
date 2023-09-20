package main

import (
	"bufio"
	"fmt"
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

	file, err := os.Open(filePath)

	converter := convert.New(1)

	error.HandleError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	theme := theme.New(scanner)

	fmt.Printf("Starting work on %s theme", theme.Name)

	converter.Convert(theme)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
