package convert

import (
	"log"

	"github.com/Inalegwu/babel/theme"
)

const (
	HelixThemeConfig = 1
)

type Converter struct {
	to int
}

func New(to int) Converter {
	return Converter{
		to,
	}
}

func (c *Converter) Convert(theme theme.Theme) error {
	if c.to == HelixThemeConfig {
		c.toHelixConfig(theme)
	} else {
		log.Fatal("Unsupported Theme config provided , If there's one you want to see , create a PR")
	}
	return nil
}

type ColorCodes []string

func (c *Converter) toHelixConfig(theme theme.Theme) {
	log.Printf("Converting %s to Helix Theme Format", theme.Name)
	log.Printf("Theme type is %s", theme.Theme_type)
	var colorCodes ColorCodes
	// var palette helix.Palette

	for _, t := range theme.Colors {
		colorCodes = append(colorCodes, t)
	}

	log.Printf("Found %v Color Codes", len(colorCodes))
}
