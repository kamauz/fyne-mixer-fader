package tapes

import (
	"fmt"

	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/models"
)

type logarithmicTape struct {
	min          float32
	max          float32
	defaultValue float32
	scaleValues  []float32
}

func (t *logarithmicTape) GetMin() float32 {
	return t.min
}

func (t *logarithmicTape) GetMax() float32 {
	return t.max
}

func (t *logarithmicTape) GetDefaultValue() float32 {
	return t.defaultValue
}

func (t *logarithmicTape) FromYPositionToValue(
	yPos,
	trackHeight,
	paddingTop float32,
) (newValue float32) {
	// Get the widget size
	if trackHeight == 0 {
		return
	}

	// Clamp Y to [0, height]
	yPos -= paddingTop
	if yPos < 0 {
		yPos = 0
	}
	if yPos > trackHeight {
		yPos = trackHeight
	}

	proportion := 1.0 - float64(yPos)/float64(trackHeight)

	// Non-linear mapping: position proportion → dB value
	// This gives finer control in lower dB ranges (typical mixer behavior)
	if proportion <= 0.25 {
		// Bottom 25% of slider: -120 to -60 dB
		newValue = t.min + float32(proportion/0.25)*(-60-t.min)
	} else if proportion <= 0.5 {
		// Next 25%: -60 to -20 dB
		newValue = -60 + float32((proportion-0.25)/0.25)*(-20-(-60))
	} else if proportion <= 0.75 {
		// Next 25%: -20 to 0 dB
		newValue = -20 + float32((proportion-0.5)/0.25)*(0-(-20))
	} else {
		// Top 25%: 0 to +12 dB (or max)
		newValue = 0 + float32((proportion-0.75)/0.25)*(t.max-0)
	}

	// Clamp value to [Min, Max]
	if newValue < t.min {
		newValue = t.min
	} else if newValue > t.max {
		newValue = t.max
	}

	return
}

func (t *logarithmicTape) FromValueToDisplayRatio(
	value float32,
) float32 {
	// from linear to logaritmic
	switch {
	case value <= -60:
		return 0.25 * (value - t.min) / (-60 - t.min)
	case value <= -20:
		return 0.25 + 0.25*(value-(-60))/(-20-(-60))
	case value <= 0:
		return 0.50 + 0.25*(value-(-20))/(0-(-20))
	default:
		return 0.75 + 0.25*(value-0)/(t.max-0)
	}
}

func (t *logarithmicTape) CalculateSideScale(
	trackHeight,
	paddingTop,
	labelWidth,
	labelHeight,
	labelX float32,
) []models.ScaleLabelSpec {

	specs := make([]models.ScaleLabelSpec, 0, len(t.scaleValues))
	for _, v := range t.scaleValues {

		// generate logarithmic ratio
		var displayRatio float32
		switch {
		case v <= -60:
			displayRatio = 0.25 * (v - (t.min)) / (-60 - (t.min))
		case v <= -20:
			displayRatio = 0.25 + 0.25*(v-(-60))/(-20-(-60))
		case v <= 0:
			displayRatio = 0.50 + 0.25*(v-(-20))/(0-(-20))
		default:
			displayRatio = 0.75 + 0.25*(v-0)/(t.max-0)
		}

		labelY := paddingTop + (1.0-displayRatio)*trackHeight - labelHeight/2

		var text string
		switch {
		case v == t.min:
			text = "-inf"
		case v == 0:
			text = "0 dB"
		default:
			text = fmt.Sprintf("%.0f", v)
		}

		specs = append(specs, models.ScaleLabelSpec{
			Text:   text,
			X:      labelX,
			Y:      labelY,
			Width:  labelWidth,
			Height: labelHeight,
		})
	}
	return specs
}

func NewLogarithmicTape() *logarithmicTape {
	return &logarithmicTape{
		min:          -120,
		max:          12,
		defaultValue: 0,
		scaleValues:  []float32{-82, -60, -30, -12, 0, 6, 12},
	}
}
