package helix

import "github.com/Inalegwu/babel/theme"

type (
	HelixThemeConfig struct{}
	Palette          map[string]string
)

func New() HelixThemeConfig {
	return HelixThemeConfig{}
}

func FromVsConfig(theme.Theme) {
}
