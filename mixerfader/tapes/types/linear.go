package tapes

import (
	"fmt"

	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/models"
)

type linearTape struct {
	min          float32
	max          float32
	defaultValue float32
	scaleValues  []float32
}

func (t *linearTape) GetMin() float32 {
	return t.min
}

func (t *linearTape) GetMax() float32 {
	return t.max
}

func (t *linearTape) GetDefaultValue() float32 {
	return t.defaultValue
}

func (t *linearTape) FromYPositionToValue(
	yPos,
	trackHeight,
	paddingTop float32,
) (newValue float32) {

	if trackHeight <= 0 {
		return t.min
	}

	// sottrai il paddingTop per lavorare nello spazio della track
	yPos -= paddingTop

	if yPos < 0 {
		yPos = 0
	} else if yPos > trackHeight {
		yPos = trackHeight
	}

	// top (0) → Max, bottom (trackHeight) → Min
	proportion := 1.0 - yPos/trackHeight

	newValue = t.min + proportion*(t.max-t.min)

	if newValue < t.min {
		newValue = t.min
	} else if newValue > t.max {
		newValue = t.max
	}

	return
}

func (t *linearTape) FromValueToDisplayRatio(
	value float32,
) float32 {
	valueRange := t.max - t.min
	if valueRange == 0 {
		return 0
	}
	ratio := (value - t.min) / valueRange
	if ratio < 0 {
		return 0
	}
	if ratio > 1 {
		return 1
	}
	return ratio
}

func (t *linearTape) CalculateSideScale(
	trackHeight,
	paddingTop,
	labelWidth,
	labelHeight,
	labelX float32,
) []models.ScaleLabelSpec {

	valueRange := t.max - t.min
	if valueRange == 0 {
		valueRange = 1
	}

	specs := make([]models.ScaleLabelSpec, 0, len(t.scaleValues))
	for _, v := range t.scaleValues {

		// generate linear ratio
		displayRatio := (v - t.min) / valueRange

		// clamp sicurezza
		if displayRatio < 0 {
			displayRatio = 0
		}
		if displayRatio > 1 {
			displayRatio = 1
		}

		// get label y position
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

func NewLinearTape() *linearTape {
	return &linearTape{
		min:          -120,
		max:          12,
		defaultValue: 0,
		scaleValues:  []float32{-96, -84, -72, -60, -48, -36, -24, -12, 0, 12},
	}
}
