package helix

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	errorCustom "github.com/Inalegwu/babel/error"
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

type ColorApiResponse struct {
	Hex  map[string]string `json:"hex"`
	Name map[string]any    `json:"name"`
}

func New(colorPalette []ColorApiResponse, theme theme.Theme) HelixThemeConfig {
	// var constant string
	// for _, v := range theme.TokenColors {
	// 	scope := v["scope"]
	// 	settings := v["settings"]

	// 	log.Printf("%s , %s", scope, settings)
	// }

	return HelixThemeConfig{
		// define the color scheme type "dark"|"light"|"other"
		"type": theme.Theme_type,
		// tree-sitter scopes
	}
}

func MakeColorPalette(colorCodes []string) []ColorApiResponse {
	log.Printf("Generating Color Palette.This may take a while")

	var colorApiResponseArray []ColorApiResponse

	for _, v := range colorCodes {
		parts := strings.Split(v, "#")
		resp, err := http.Get("https://www.thecolorapi.com/id?hex=" + parts[1])

		var response ColorApiResponse

		errorCustom.HandleError(err)

		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &response)

		errorCustom.HandleError(err)

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
