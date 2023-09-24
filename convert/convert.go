package convert

import (
	"log"

	"github.com/Inalegwu/babel/adapters/helix"
	customError "github.com/Inalegwu/babel/error"
	"github.com/Inalegwu/babel/palette"
	"github.com/Inalegwu/babel/theme"
	"github.com/Inalegwu/babel/utils"
)

const (
	HelixThemeConfig = 1
	VimThemeConfig   = 2
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
	} else if c.to == VimThemeConfig {
		log.Fatal("[WIP] Vim theme config adapter still under implementation")
	} else {
		log.Fatal("Unsupported Theme config provided , If there's one you want to see , create a PR")
	}
	return nil
}

type ColorCodes []string

func (c *Converter) toHelixConfig(theme theme.Theme) {
	log.Printf("Converting %s to Theme Format", theme.Name)
	log.Printf("Theme type is %s", theme.Theme_type)

	uniqueColorCodes := findUniqueColorCodes(theme)

	log.Printf("Found %v unique color codes", len(uniqueColorCodes))

	palette := palette.MakeColorPalette(uniqueColorCodes)

	helixTheme := helix.New(palette, theme)

	toml, err := helix.WriteToml(helixTheme)

	log.Printf("%v", string(toml))

	customError.HandleError(err)
}

func (c *Converter) toVimConfig(theme theme.Theme) {
	log.Printf("Converting %s to Theme Format", theme.Name)
	log.Printf("Theme type is %s", theme.Theme_type)

	uniqueColorCodes := findUniqueColorCodes(theme)

	log.Printf("Found %v unique color codes", len(uniqueColorCodes))
}

func findUniqueColorCodes(theme theme.Theme) ColorCodes {
	var colorCodes ColorCodes

	for _, t := range theme.Colors {
		colorCodes = append(colorCodes, t)
	}

	var uniqueColorCodes ColorCodes

	for _, t := range colorCodes {
		alreadyInUnique := utils.Find(uniqueColorCodes, t)
		if !alreadyInUnique {
			uniqueColorCodes = append(uniqueColorCodes, t)
		}
	}

	return uniqueColorCodes
}
