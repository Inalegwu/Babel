package palette

import (
	"log"
	"strings"

	"github.com/Inalegwu/babel/request"
)

func MakeColorPalette(colorCodes []string) []request.ColorApiResponse {
	log.Printf("Generating Color Palette.This may take a while...")

	var colorApiResponseArray []request.ColorApiResponse

	for _, v := range colorCodes {
		parts := strings.Split(v, "#")

		response := request.GetColorCodeName(parts[1])

		colorApiResponseArray = append(colorApiResponseArray, response)
	}

	log.Printf("Color Palette Generated Successfully")

	return colorApiResponseArray
}
