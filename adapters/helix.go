package adapters

import "github.com/Inalegwu/babel/theme"

type HelixThemeConfig struct{}

func New() HelixThemeConfig {
	return HelixThemeConfig{}
}

func FromVsConfig(theme.Theme) {
}
