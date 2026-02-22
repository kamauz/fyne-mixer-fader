package config

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type MiddleLineStyle struct {
	StrokeColor color.NRGBA
	StrokeWidth float32
}

type LabelBoxStyle struct {
	Color        color.NRGBA
	CornerRadius float32
}

type LabelStyle struct {
	Color     color.NRGBA
	TextAlign fyne.TextAlign
	Bold      bool
}

type ThumbStyle struct {
	// color
	Color       color.NRGBA
	BorderColor color.NRGBA

	StrokeWidth  float32
	CornerRadius float32
	Bold         bool

	TextAlign fyne.TextAlign

	// show / hide
	ShowLabel    bool
	ShowLabelBox bool

	Width  float32
	Height float32

	// children
	MiddleLineStyle MiddleLineStyle
	LabelBoxStyle   LabelBoxStyle
	LabelStyle      LabelStyle
}
