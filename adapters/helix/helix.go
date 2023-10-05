package helix

import (
	"log"

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

type ScopeSettings struct {
	scope    Scope
	settings Settings
}

func New(colorPalette []request.ColorApiResponse, theme theme.Theme) HelixThemeConfig {
	log.Print("Using generated color palette to populate Helix Theme Config")

	for _, t := range theme.TokenColors {
		// scope := t["scope"]
		var settings interface{} = t["settings"]

		log.Printf("%v", settings)

		v, err := settings.(Settings)

		if err {
			log.Fatal("Invalid type assertion")
		}

		log.Printf("%v", v)

	}

	return HelixThemeConfig{
		// define the color scheme type "dark"|"light"|"other"
		"type": theme.Theme_type,
		// tree-sitter scopes
		"palette": colorPalette,
	}
}

func WriteToml(helixThemeConfig HelixThemeConfig, theme theme.Theme) ([]byte, error) {
	log.Printf("Writing Helix Theme Config to %s.toml", theme.Name)
	toml, err := toml.Marshal(helixThemeConfig)
	if err != nil {
		return nil, err
	}
	return toml, nil
}
