package config

import "image/color"

type SideScaleStyle struct {
	Color    color.NRGBA
	TextSize float32
}

type ShadowStyle struct {
	Color  color.NRGBA
	Offset float32
}

type TrackStyle struct {

	// colors
	Color           color.NRGBA
	ZeroDBLineColor color.NRGBA

	// shadows
	ShadowColor  color.NRGBA
	ShadowOffset float32

	// sizes
	Width  float32
	Radius float32

	// label sizes
	LabelWidth  float32
	LabelHeight float32

	// children
	SideScaleStyle SideScaleStyle
	ShadowStyle    ShadowStyle
}
