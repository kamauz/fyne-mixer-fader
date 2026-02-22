package theme

import (
	config "github.com/kamauz/fyne-mixer-fader/mixerfader/theme/config"
)

// Theme represents the theme for the mixer fader widget.
type Theme struct {
	config *config.ThemeConfig
}

// NewTheme creates and returns a new Theme instance with the provided configuration.
func NewTheme(
	config *config.ThemeConfig,
) *Theme {
	return &Theme{
		// config
		config: config,
	}
}
