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

type ColorApiResponse struct {
	Hex  map[string]string `json:"hex"`
	Name map[string]string `json:"name"`
	Rgb  map[string]string `json:"rgb"`
}

func New() HelixThemeConfig {
	return HelixThemeConfig{}
}

func MakeColorPalette(colorCodes []string, theme theme.Theme) {
	log.Printf("Generating Color Palette")

	var colorApiResponse []ColorApiResponse

	// for _, v := range theme.Colors {
	// 	parts := strings.Split(v, "#")
	// 	resp, err := http.Get("https://www.thecolorapi.com/id?" + parts[1])

	// 	error.HandleError(err)

	// 	err = json.Unmarshal([]byte(resp), &colorApiResponse)

	// 	error.HandleError(err)

	// 	log.Printf("%v", resp)

	// }

	badge := strings.Split(theme.Colors["badge.background"], "#")

	resp, err := http.Get("https://www.thecolorapi.com/id?hex=" + badge[1])

	error.HandleError(err)

	body, err := ioutil.ReadAll(resp.Body)

	error.HandleError(err)

	err = json.Unmarshal(body, &colorApiResponse)

	log.Printf("%v", colorApiResponse)
}

func FromVsConfig(theme.Theme) {
}
