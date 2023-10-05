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
	Settings map[string]any
)

type ScopeSettings struct {
	scope    Scope
	settings Settings
}

func New(colorPalette []request.ColorApiResponse, theme theme.Theme) HelixThemeConfig {
	log.Print("Using generated color palette to populate Helix Theme Config")

	for _, t := range theme.TokenColors {
		settings := t["settings"]
		scopes := t["scopes"]
		log.Printf("%v,%v", settings, scopes)

		var settingsInterface interface{} = settings.(Settings)

		log.Printf("%v", settingsInterface)

	}

	log.Printf("%v", theme.TokenColors)

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
