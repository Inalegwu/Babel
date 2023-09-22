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

type Name struct {
	Value           string `json:"value"`
	ClosestNamedHex string `json:"closest_named_hex"`
	ExactMatchName  bool   `json:"exact_match_name"`
	Distance        int    `json:"distance"`
}

func New(colorPalette []request.ColorApiResponse, theme theme.Theme) HelixThemeConfig {
	// convert color palette to use
	// in theme config definition

	// var palette map[string]string

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
