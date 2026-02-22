package mixerfader

import "github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/models"

func (m *MixerFader) GetTapeMinValue() float32 {
	return m.tape.GetMin()
}

func (m *MixerFader) FromValueToTapeDisplayRatio(value float32) float32 {
	return m.tape.FromValueToDisplayRatio(value)
}

func (m *MixerFader) CalculateTapeSideScale(
	trackHeight,
	paddingTop,
	labelWidth,
	labelHeight,
	labelX float32) []models.ScaleLabelSpec {
	return m.tape.CalculateSideScale(
		trackHeight,
		paddingTop,
		labelWidth,
		labelHeight,
		labelX,
	)
}
