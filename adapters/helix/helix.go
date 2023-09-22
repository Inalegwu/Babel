package helix

import (
	"log"
	"strings"

	"github.com/Inalegwu/babel/request"
	"github.com/Inalegwu/babel/theme"
	"github.com/pelletier/go-toml"
)

type (
	HelixThemeConfig map[string]any
	Palette          map[string]string
)

type (
	Scope    []string
	Settings map[string]string
)

func New(colorPalette []request.ColorApiResponse, theme theme.Theme) HelixThemeConfig {
	return HelixThemeConfig{
		// define the color scheme type "dark"|"light"|"other"
		"type": theme.Theme_type,
		// tree-sitter scopes

	}
}

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

func WriteToml(helixThemeConfig HelixThemeConfig) ([]byte, error) {
	toml, err := toml.Marshal(helixThemeConfig)
	if err != nil {
		return nil, err
	}
	return toml, nil
}
