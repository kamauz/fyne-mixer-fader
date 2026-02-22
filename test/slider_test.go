package test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	mixerfader "github.com/kamauz/fyne-mixer-fader/mixerfader/slider"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/enums"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/theme/config"
)

// newTestFader crea un fader con configurazione minimale.
// test.NewApp() inizializza un driver headless che evita panic da Fyne.
func newTestFader(initialValue float64, tapeType enums.TapeTypeEnum) *mixerfader.MixerFader {
	test.NewApp()
	return mixerfader.New(initialValue, tapeType, config.ThemeConfig{
		PaddingTop:    10,
		PaddingBottom: 10,
		PaddingLeft:   30,
		TrackConfig:   config.TrackStyle{Width: 7, LabelWidth: 50, LabelHeight: 18},
		ThumbConfig:   config.ThumbStyle{Width: 32, Height: 46},
	}, nil)
}

// ─── DB / SetDB ───────────────────────────────────────────────────────────────

func TestMixerFader_DB_ReturnsInitialValue(t *testing.T) {
	f := newTestFader(-12, enums.TapeTypeLinear)
	assert.Equal(t, -12.0, f.DB())
}

func TestMixerFader_SetDB_UpdatesValue(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	f.SetDB(-24)
	assert.Equal(t, -24.0, f.DB())
}

func TestMixerFader_SetDB_SameValueDoesNotFireCallback(t *testing.T) {
	called := 0
	test.NewApp()
	f := mixerfader.New(0, enums.TapeTypeLinear, config.ThemeConfig{
		TrackConfig: config.TrackStyle{Width: 7, LabelWidth: 50, LabelHeight: 18},
		ThumbConfig: config.ThumbStyle{Width: 32, Height: 46},
	}, func(v float64) { called++ })

	f.SetDB(0) // stesso valore iniziale → non deve chiamare il callback
	assert.Equal(t, 0, called)
}

func TestMixerFader_SetDB_DifferentValueFiresCallback(t *testing.T) {
	var received float64
	test.NewApp()
	f := mixerfader.New(0, enums.TapeTypeLinear, config.ThemeConfig{
		TrackConfig: config.TrackStyle{Width: 7, LabelWidth: 50, LabelHeight: 18},
		ThumbConfig: config.ThumbStyle{Width: 32, Height: 46},
	}, func(v float64) { received = v })

	f.SetDB(-6)
	assert.Equal(t, -6.0, received)
}

func TestMixerFader_SetDB_NilCallbackDoesNotPanic(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	f.OnChanged = nil
	assert.NotPanics(t, func() { f.SetDB(-12) })
}

func TestMixerFader_SetDB_MultipleUpdates(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	for _, v := range []float64{-12, -6, 0, 6} {
		f.SetDB(v)
		assert.Equal(t, v, f.DB())
	}
}

// ─── GetTapeMinValue ──────────────────────────────────────────────────────────

func TestMixerFader_GetTapeMinValue_Linear(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	assert.Equal(t, float32(-120), f.GetTapeMinValue())
}

func TestMixerFader_GetTapeMinValue_Logarithmic(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLogarithmic)
	assert.Equal(t, float32(-120), f.GetTapeMinValue())
}

// ─── FromValueToTapeDisplayRatio ─────────────────────────────────────────────

func TestMixerFader_DisplayRatio_ZeroDBLinear(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	// 0 dB lineare: (0 - (-120)) / (12 - (-120)) = 120/132
	expected := float32(120.0 / 132.0)
	assert.InDelta(t, expected, f.FromValueToTapeDisplayRatio(0), 1e-4)
}

func TestMixerFader_DisplayRatio_ZeroDBLogarithmic(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLogarithmic)
	// 0 dB logaritmico è al 75° percentile per definizione
	assert.InDelta(t, 0.75, f.FromValueToTapeDisplayRatio(0), 1e-4)
}

func TestMixerFader_DisplayRatio_MinGivesZero(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	assert.InDelta(t, 0.0, f.FromValueToTapeDisplayRatio(-120), 1e-5)
}

func TestMixerFader_DisplayRatio_MaxGivesOne(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	assert.InDelta(t, 1.0, f.FromValueToTapeDisplayRatio(12), 1e-5)
}

// ─── CalculateTapeSideScale ───────────────────────────────────────────────────

func TestMixerFader_SideScale_LinearReturnsLabels(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	specs := f.CalculateTapeSideScale(300, 10, 50, 18, 0)
	assert.NotEmpty(t, specs)
}

func TestMixerFader_SideScale_LogarithmicReturnsLabels(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLogarithmic)
	specs := f.CalculateTapeSideScale(300, 10, 50, 18, 0)
	assert.NotEmpty(t, specs)
}

func TestMixerFader_SideScale_HasZeroDBLabel(t *testing.T) {
	for _, tapeType := range []enums.TapeTypeEnum{enums.TapeTypeLinear, enums.TapeTypeLogarithmic} {
		t.Run(string(tapeType), func(t *testing.T) {
			f := newTestFader(0, tapeType)
			specs := f.CalculateTapeSideScale(300, 10, 50, 18, 0)
			found := false
			for _, s := range specs {
				if s.Text == "0 dB" {
					found = true
					break
				}
			}
			assert.True(t, found, "deve esserci un label '0 dB' per tape %s", tapeType)
		})
	}
}

func TestMixerFader_SideScale_LabelXMatchesInput(t *testing.T) {
	f := newTestFader(0, enums.TapeTypeLinear)
	labelX := float32(5)
	specs := f.CalculateTapeSideScale(300, 10, 50, 18, labelX)
	for _, s := range specs {
		assert.Equal(t, labelX, s.X)
	}
}

// ─── New — panic su tape type sconosciuto ─────────────────────────────────────

func TestNew_PanicsOnUnknownTapeType(t *testing.T) {
	test.NewApp()
	assert.Panics(t, func() {
		mixerfader.New(0, enums.TapeTypeEnum("INVALID"), config.ThemeConfig{
			TrackConfig: config.TrackStyle{Width: 7, LabelWidth: 50, LabelHeight: 18},
			ThumbConfig: config.ThumbStyle{Width: 32, Height: 46},
		}, nil)
	})
}

func TestNew_InitialValueIsPreserved(t *testing.T) {
	f := newTestFader(-45.3, enums.TapeTypeLogarithmic)
	assert.Equal(t, -45.3, f.DB())
}
