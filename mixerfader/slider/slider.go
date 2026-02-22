package mixerfader

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/enums"
	tapeTypes "github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/types"
	fadertheme "github.com/kamauz/fyne-mixer-fader/mixerfader/theme"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/theme/config"
)

// MixerFader is a custom vertical slider
type MixerFader struct {
	widget.BaseWidget
	value float64
	tape  tapeTypes.Tapable

	// styling
	theme *fadertheme.Theme

	// events
	hovered bool

	// callbacks
	OnChanged func(float64)
}

func (s *MixerFader) SetDB(v float64) {
	if s.value == v {
		return
	}
	s.value = v
	if s.OnChanged != nil {
		s.OnChanged(v)
	}
	s.Refresh()
}

func (s *MixerFader) DB() float64 {
	return s.value
}

// New creates and returns a new MixerFader widget.
func New(
	value float64,

	// tape type ("logarithmic", linear, ...)
	tapeType enums.TapeTypeEnum,

	// styling
	themeConfig config.ThemeConfig,

	// callbacks
	onChanged func(float64)) *MixerFader {

	// init theme from theme config
	theme := fadertheme.NewTheme(
		&themeConfig,
	)

	// define tape with factory
	tape, err := tapes.CreateTapeFactory(tapeType)
	if err != nil {
		panic(err) // oppure fallback default
	}

	// init mixer fader extending fyne slider
	s := &MixerFader{
		value:     value,
		theme:     theme,
		tape:      tape,
		OnChanged: onChanged,
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *MixerFader) CreateRenderer() fyne.WidgetRenderer {

	return newMixerFaderRenderer(
		// slider
		s,

		// theme
		s.theme,
	)
}
