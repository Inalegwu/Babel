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
	var scopes []ScopeSettings

	// still not sure if this works
	// must test
	for _, t := range theme.TokenColors {
		var scope interface{} = t["scope"]
		var settings interface{} = t["settings"]

		settingsVal, err := settings.(Settings)

		if !err {
			log.Fatal("Invalid type coercion on Settings")
		}

		scopeVal, err := scope.(Scope)

		if !err {
			log.Fatal("Invalid type coercion on Scope")
		}

		scopeSetting := ScopeSettings{
			scope:    scopeVal,
			settings: settingsVal,
		}

		scopes = append(scopes, scopeSetting)

	}

	log.Printf("%v", scopes)

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
