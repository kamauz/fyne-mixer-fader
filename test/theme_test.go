package test

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/theme/config"
	"github.com/stretchr/testify/assert"
)

// ─── ThemeConfig ──────────────────────────────────────────────────────────────

func TestThemeConfig_ZeroValue(t *testing.T) {
	cfg := config.ThemeConfig{}
	assert.Equal(t, float32(0), cfg.PaddingTop)
	assert.Equal(t, float32(0), cfg.PaddingRight)
	assert.Equal(t, float32(0), cfg.PaddingBottom)
	assert.Equal(t, float32(0), cfg.PaddingLeft)
}

func TestThemeConfig_PaddingsAreIndependent(t *testing.T) {
	cfg := config.ThemeConfig{
		PaddingTop:    10,
		PaddingRight:  5,
		PaddingBottom: 15,
		PaddingLeft:   30,
	}
	assert.Equal(t, float32(10), cfg.PaddingTop)
	assert.Equal(t, float32(5), cfg.PaddingRight)
	assert.Equal(t, float32(15), cfg.PaddingBottom)
	assert.Equal(t, float32(30), cfg.PaddingLeft)
}

func TestThemeConfig_HoldsThumbAndTrackConfig(t *testing.T) {
	cfg := config.ThemeConfig{
		ThumbConfig: config.ThumbStyle{Width: 32, Height: 46},
		TrackConfig: config.TrackStyle{Width: 8},
	}
	assert.Equal(t, float32(32), cfg.ThumbConfig.Width)
	assert.Equal(t, float32(46), cfg.ThumbConfig.Height)
	assert.Equal(t, float32(8), cfg.TrackConfig.Width)
}

func TestThemeConfig_TotalHorizontalPadding(t *testing.T) {
	cfg := config.ThemeConfig{PaddingLeft: 30, PaddingRight: 5}
	total := cfg.PaddingLeft + cfg.PaddingRight
	assert.Equal(t, float32(35), total)
}

func TestThemeConfig_TotalVerticalPadding(t *testing.T) {
	cfg := config.ThemeConfig{PaddingTop: 10, PaddingBottom: 10}
	total := cfg.PaddingTop + cfg.PaddingBottom
	assert.Equal(t, float32(20), total)
}

// ─── TrackStyle ───────────────────────────────────────────────────────────────

func TestTrackStyle_ZeroValue(t *testing.T) {
	ts := config.TrackStyle{}
	assert.Equal(t, float32(0), ts.Width)
	assert.Equal(t, float32(0), ts.Radius)
	assert.Equal(t, float32(0), ts.LabelWidth)
	assert.Equal(t, float32(0), ts.LabelHeight)
}

func TestTrackStyle_ColorsAreIndependent(t *testing.T) {
	ts := config.TrackStyle{
		Color:           color.NRGBA{R: 22, G: 22, B: 22, A: 255},
		ZeroDBLineColor: color.NRGBA{R: 255, G: 200, B: 0, A: 180},
	}
	assert.Equal(t, uint8(22), ts.Color.R)
	assert.Equal(t, uint8(255), ts.ZeroDBLineColor.R)
	assert.NotEqual(t, ts.Color, ts.ZeroDBLineColor)
}

func TestTrackStyle_ShadowStyleEmbedded(t *testing.T) {
	ts := config.TrackStyle{
		ShadowStyle: config.ShadowStyle{
			Color:  color.NRGBA{R: 200, G: 200, B: 200, A: 140},
			Offset: 3,
		},
	}
	assert.Equal(t, float32(3), ts.ShadowStyle.Offset)
	assert.Equal(t, uint8(200), ts.ShadowStyle.Color.R)
}

func TestTrackStyle_SideScaleStyleEmbedded(t *testing.T) {
	ts := config.TrackStyle{
		SideScaleStyle: config.SideScaleStyle{
			Color:    color.NRGBA{R: 250, G: 250, B: 250, A: 255},
			TextSize: 10,
		},
	}
	assert.Equal(t, float32(10), ts.SideScaleStyle.TextSize)
}

func TestTrackStyle_LabelDimensionsCanBeSet(t *testing.T) {
	ts := config.TrackStyle{LabelWidth: 50, LabelHeight: 18}
	assert.Equal(t, float32(50), ts.LabelWidth)
	assert.Equal(t, float32(18), ts.LabelHeight)
}

// ─── ShadowStyle ─────────────────────────────────────────────────────────────

func TestShadowStyle_ZeroValue(t *testing.T) {
	s := config.ShadowStyle{}
	assert.Equal(t, float32(0), s.Offset)
	assert.Equal(t, color.NRGBA{}, s.Color)
}

func TestShadowStyle_PositiveOffset(t *testing.T) {
	s := config.ShadowStyle{Offset: 4}
	assert.Greater(t, s.Offset, float32(0))
}

// ─── SideScaleStyle ───────────────────────────────────────────────────────────

func TestSideScaleStyle_ZeroValue(t *testing.T) {
	s := config.SideScaleStyle{}
	assert.Equal(t, float32(0), s.TextSize)
}

func TestSideScaleStyle_TextSizeCanBeSet(t *testing.T) {
	s := config.SideScaleStyle{TextSize: 10}
	assert.Equal(t, float32(10), s.TextSize)
}

// ─── ThumbStyle ───────────────────────────────────────────────────────────────

func TestThumbStyle_ZeroValue(t *testing.T) {
	ts := config.ThumbStyle{}
	assert.Equal(t, float32(0), ts.Width)
	assert.Equal(t, float32(0), ts.Height)
	assert.False(t, ts.ShowLabel)
	assert.False(t, ts.ShowLabelBox)
	assert.False(t, ts.Bold)
}

func TestThumbStyle_ShowFlagsAreIndependent(t *testing.T) {
	ts := config.ThumbStyle{ShowLabel: true, ShowLabelBox: false}
	assert.True(t, ts.ShowLabel)
	assert.False(t, ts.ShowLabelBox)
}

func TestThumbStyle_DimensionsCanBeSet(t *testing.T) {
	ts := config.ThumbStyle{Width: 32, Height: 46}
	assert.Equal(t, float32(32), ts.Width)
	assert.Equal(t, float32(46), ts.Height)
}

func TestThumbStyle_ColorsAreIndependent(t *testing.T) {
	ts := config.ThumbStyle{
		Color:       color.NRGBA{R: 255, G: 255, B: 255, A: 220},
		BorderColor: color.NRGBA{R: 200, G: 200, B: 200, A: 190},
	}
	assert.NotEqual(t, ts.Color, ts.BorderColor)
}

func TestThumbStyle_ChildStylesEmbedded(t *testing.T) {
	ts := config.ThumbStyle{
		MiddleLineStyle: config.MiddleLineStyle{StrokeWidth: 2},
		LabelBoxStyle:   config.LabelBoxStyle{CornerRadius: 6},
		LabelStyle:      config.LabelStyle{Bold: true, TextAlign: fyne.TextAlignCenter},
	}
	assert.Equal(t, float32(2), ts.MiddleLineStyle.StrokeWidth)
	assert.Equal(t, float32(6), ts.LabelBoxStyle.CornerRadius)
	assert.True(t, ts.LabelStyle.Bold)
	assert.Equal(t, fyne.TextAlignCenter, ts.LabelStyle.TextAlign)
}

// ─── MiddleLineStyle ─────────────────────────────────────────────────────────

func TestMiddleLineStyle_ZeroValue(t *testing.T) {
	m := config.MiddleLineStyle{}
	assert.Equal(t, float32(0), m.StrokeWidth)
	assert.Equal(t, color.NRGBA{}, m.StrokeColor)
}

func TestMiddleLineStyle_CanBePopulated(t *testing.T) {
	m := config.MiddleLineStyle{
		StrokeColor: color.NRGBA{R: 0, G: 0, B: 0, A: 255},
		StrokeWidth: 2,
	}
	assert.Equal(t, float32(2), m.StrokeWidth)
	assert.Equal(t, uint8(255), m.StrokeColor.A)
}

// ─── LabelBoxStyle ────────────────────────────────────────────────────────────

func TestLabelBoxStyle_ZeroValue(t *testing.T) {
	l := config.LabelBoxStyle{}
	assert.Equal(t, float32(0), l.CornerRadius)
}

func TestLabelBoxStyle_CornerRadiusCanBeSet(t *testing.T) {
	l := config.LabelBoxStyle{CornerRadius: 6}
	assert.Equal(t, float32(6), l.CornerRadius)
}

// ─── LabelStyle ──────────────────────────────────────────────────────────────

func TestLabelStyle_ZeroValue(t *testing.T) {
	l := config.LabelStyle{}
	assert.False(t, l.Bold)
	assert.Equal(t, fyne.TextAlign(0), l.TextAlign)
}

func TestLabelStyle_BoldAndAlignCanBeSet(t *testing.T) {
	l := config.LabelStyle{Bold: true, TextAlign: fyne.TextAlignCenter}
	assert.True(t, l.Bold)
	assert.Equal(t, fyne.TextAlignCenter, l.TextAlign)
}

func TestLabelStyle_ColorCanBeSet(t *testing.T) {
	c := color.NRGBA{R: 250, G: 250, B: 250, A: 240}
	l := config.LabelStyle{Color: c}
	assert.Equal(t, c, l.Color)
}
