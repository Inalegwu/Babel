package helix

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Inalegwu/babel/error"
	"github.com/Inalegwu/babel/theme"
)

type (
	HelixThemeConfig struct{}
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
	Name map[string]Name   `json:"name"`
}

func New() HelixThemeConfig {
	return HelixThemeConfig{}
}

func MakeColorPalette(colorCodes []string, theme theme.Theme) {
	log.Printf("Generating Color Palette")

	// var colorApiResponse ColorApiResponse

	// var colorApiResponseArray []ColorApiResponse

	for _, v := range theme.Colors {
		parts := strings.Split(v, "#")
		resp, err := http.Get("https://www.thecolorapi.com/id?hex=" + parts[1])

		var response ColorApiResponse

		error.HandleError(err)

		body, err := ioutil.ReadAll(resp.Body)

		log.Printf("%s", string(body))

		err = json.Unmarshal(body, &response)

		error.HandleError(err)

		log.Printf("%v", response)

	}
}

func FromVsConfig(theme.Theme) {
}
