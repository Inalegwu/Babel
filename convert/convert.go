package convert

import (
	"log"
	"strings"

	"github.com/Inalegwu/vsconvert/theme"
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

	var colorCodes ColorCodes

	for _, t := range theme.Colors {
		parts := strings.Split(t, "#")
		// splitting by space because strings.Split(t,"#") strips the #
		// I forgot this 45 Seconds after implementing this
		// this journey is about to be wild
		if parts[0] == "" {
			colorCodes = append(colorCodes, t)
		}
	}
	log.Printf("Found %v Color Codes", len(colorCodes))
}
