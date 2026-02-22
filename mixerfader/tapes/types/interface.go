package tapes

import "github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/models"

type Tapable interface {
	GetMin() float32
	GetMax() float32
	GetDefaultValue() float32
	FromYPositionToValue(yPos, trackHeight, paddingTop float32) float32
	FromValueToDisplayRatio(value float32) float32
	CalculateSideScale(trackHeight, paddingTop, labelWidth, labelHeight, labelX float32) []models.ScaleLabelSpec
}
