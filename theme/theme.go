package theme

import (
	"bufio"
	"encoding/json"
	"log"

	"github.com/Inalegwu/vsconvert/error"
)

type ActivityBarConfig struct {
	activeBorder       string
	background         string
	dropBorder         string
	foreground         string
	inactiveForeground string
}

type ActivityBarBadgeConfig struct {
	background string
	foreground string
}

type BadgeConfig struct {
	background string
	foreground string
}

type BannerConfig struct {
	background     string
	foreground     string
	iconForeground string
}

type BreadCrumbConfig struct {
	activeSelectionForeground string
	background                string
	focusForeground           string
	foreground                string
}

type BreadCrumbPickerConfig struct {
	background string
}

type ButtonConfig struct {
	background               string
	foreground               string
	hoverBackground          string
	secondaryBackground      string
	secondaryForeground      string
	secondaryHoverBackground string
}

type ChartsConfig struct {
	lines string
}

type Colors struct {
	activityBar      ActivityBarConfig
	activityBarBadge ActivityBarBadgeConfig
	badge            BadgeConfig
	banner           BannerConfig
	breadcrumb       BreadCrumbConfig
	breadcrumbPicker BreadCrumbPickerConfig
	button           ButtonConfig
	charts           ChartsConfig
}

type Theme struct {
	Name       string
	theme_type string
	colors     Colors
}

func New(scanner *bufio.Scanner) Theme {
	log.Print("Reading JSON Theme Config")
	var themeConfig Theme

	scanner.Split(bufio.ScanLines)

	err := json.Unmarshal(scanner.Bytes(), &themeConfig)

	error.HandleError(err)

	return themeConfig
}
