package adapters

import "github.com/Inalegwu/vsconvert/theme"

type HelixThemeConfig struct{}

func New() HelixThemeConfig {
	return HelixThemeConfig{}
}

func FromVsConfig(theme.Theme) {
}
