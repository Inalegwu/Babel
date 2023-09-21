package convert

import (
	"log"
	// "github.com/Inalegwu/babel/adapters/helix"
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

	var colorCodes ColorCodes
	// var palette helix.Palette

	for _, t := range theme.Colors {
		colorCodes = append(colorCodes, t)
	}

	var uniqueColorCodes []string

	for i := 0; i < len(colorCodes)-1; i++ {
		exists := utils.Find(colorCodes, colorCodes[i])
		if !exists {
			uniqueColorCodes = append(uniqueColorCodes, colorCodes[i])
		}
	}

	log.Printf("Found %v Color Codes", len(colorCodes))
	log.Printf("Searching for duplicates...")
	if len(uniqueColorCodes) == 0 {
		log.Print("No Duplicates found")
	} else {
		log.Printf("%v Duplicates found", len(uniqueColorCodes))
	}
}
