package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/Inalegwu/babel/convert"
	customError "github.com/Inalegwu/babel/error"
	"github.com/Inalegwu/babel/theme"
)

func main() {
	log.SetPrefix("BABEL :: ")
	filePath := os.Getenv("FILE")
	convertFlag := os.Getenv("FLAG")

	// ensure appropriate environment flags are passed
	// should probably turn this into a CLI app eventually just
	// because
	if filePath == "" {
		log.Fatal("Please provide a file path by passing FILE=`file_path` ")
	}

	if convertFlag == "" {
		log.Fatal("Please provide a convert flag by passing FLAG=`flag`(int)")
	}

	// convert flag value passed from environment to
	// int value
	flagInt, err := strconv.Atoi(convertFlag)

	customError.HandleError(err)

	// try opening file
	file, err := os.Open(filePath)

	customError.HandleError(err)

	// initialize the converter with a ConverterType flag (int)
	converter := convert.New(flagInt)

	// get file info to use for unmarshalling
	fileStat, err := file.Stat()
	customError.HandleError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	theme := theme.New(scanner, fileStat)

	converter.Convert(theme)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
