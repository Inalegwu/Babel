package convert

import (
	"log"

	"github.com/Inalegwu/babel/adapters/helix"
	customError "github.com/Inalegwu/babel/error"
	"github.com/Inalegwu/babel/theme"
	"github.com/Inalegwu/babel/utils"
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

	for _, t := range theme.Colors {
		colorCodes = append(colorCodes, t)
	}

	log.Printf("Found %v Color Codes", len(colorCodes))

	var uniqueColorCodes ColorCodes

	log.Print("Searching for Duplicates")

	for _, t := range colorCodes {
		alreadyInUnique := utils.Find(uniqueColorCodes, t)
		if !alreadyInUnique {
			uniqueColorCodes = append(uniqueColorCodes, t)
		}
	}

	log.Printf("Found %v unique color codes. These will be used to generate the Color Palette", len(uniqueColorCodes))

	palette := helix.MakeColorPalette(uniqueColorCodes)

	helixTheme := helix.New(palette, theme)

	toml, err := helix.WriteToml(helixTheme)

	customError.HandleError(err)

	log.Printf("%v", string(toml))

	log.Printf("%v", len(palette))
}
