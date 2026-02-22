package theme

import (
	"fyne.io/fyne/v2"
)

func (m *Theme) GetMinSize() (w, h float32) {
	thumbWidth := m.config.ThumbConfig.Width
	scaleWidth := m.config.PaddingLeft
	w = scaleWidth + thumbWidth + m.config.PaddingRight
	h = 220 + m.config.PaddingTop + m.config.PaddingBottom

	return
}

func (m *Theme) GetPaddings() (
	paddingTop,
	paddingRight,
	paddingBottom,
	paddingLeft float32) {

	paddingTop = m.config.PaddingTop
	paddingRight = m.config.PaddingRight
	paddingBottom = m.config.PaddingBottom
	paddingLeft = m.config.PaddingLeft

	return
}

// thumb

func (m *Theme) IsShowLabelBox() bool {
	isShowLabelBox := m.config.ThumbConfig.ShowLabelBox
	return isShowLabelBox
}

func (m *Theme) IsShowLabel() bool {
	isShowLabel := m.config.ThumbConfig.ShowLabel
	return isShowLabel
}

func (m *Theme) GetThumbPosition(
	displayRatio float32,
	thumbYAnchor *float32,
	thumbSize fyne.Size,
	size fyne.Size,
) float32 {
	thumbWidth := thumbSize.Width
	thumbHeight := thumbSize.Height
	trackHeight := size.Height - m.config.PaddingTop - m.config.PaddingBottom

	indicatorY := m.config.PaddingTop + (1.0-displayRatio)*trackHeight
	*thumbYAnchor = indicatorY - thumbHeight/2

	availableWidth := size.Width - m.config.PaddingLeft - m.config.PaddingRight
	thumbX := m.config.PaddingLeft + (availableWidth-thumbWidth)/2

	return thumbX
}

func (m *Theme) GetThumbMiddleLinePosition(
	posY float32,
	thumbSize fyne.Size,
) (lineStartX, lineEndX, lineY float32) {
	// Draw black horizontal line in the middle of the thumb
	thumbCenterX := m.config.PaddingLeft + thumbSize.Width/2
	thumbCenterY := posY + thumbSize.Height/2
	lineLength := thumbSize.Width * 0.6

	lineStartX = thumbCenterX - lineLength/2
	lineEndX = thumbCenterX + lineLength/2
	lineY = thumbCenterY

	return
}

func (m *Theme) GetThumbLabelPosition(
	size fyne.Size,
) (labelX, labelWidth float32) {
	labelWidth = m.config.TrackConfig.LabelWidth

	// center the label in the remaining space after paddingLeft and paddingRight
	availableWidth := size.Width - m.config.PaddingLeft - m.config.PaddingRight
	labelX = m.config.PaddingLeft + (availableWidth-labelWidth)/2

	return
}

func (m *Theme) GetThumbLabelBoxPosition(thumbYAnchor *float32, size fyne.Size) (w, h, x, y float32) {
	labelWidth := m.config.TrackConfig.LabelWidth

	// center the label box in the remaining space after paddingLeft and paddingRight
	availableWidth := size.Width - m.config.PaddingLeft - m.config.PaddingRight
	labelX := m.config.PaddingLeft + (availableWidth-labelWidth)/2

	w = labelWidth
	h = float32(20)
	x = labelX
	y = *thumbYAnchor - 28

	return
}

// track

func (m *Theme) GetTrackHeight(value float32) float32 {
	return value - m.config.PaddingTop - m.config.PaddingBottom
}

func (m *Theme) GetTrackLabelSize() (float32, float32) {
	trackConfig := m.config.TrackConfig
	return trackConfig.LabelWidth, trackConfig.LabelHeight
}

func (m *Theme) GetTrackPosition(
	size fyne.Size,
) (w, h, x, y float32) {
	w = m.config.TrackConfig.Width
	h = size.Height - m.config.PaddingTop - m.config.PaddingBottom

	// Center the track in the remaining space after paddingLeft and paddingRight
	availableWidth := size.Width - m.config.PaddingLeft - m.config.PaddingRight

	x = m.config.PaddingLeft + (availableWidth-m.config.TrackConfig.Width)/2
	y = m.config.PaddingTop
	return
}

func (m *Theme) GetTrackShadow(
	size fyne.Size,
) (w, h, x, y float32) {
	trackHeight := size.Height - m.config.PaddingTop - m.config.PaddingBottom

	availableWidth := size.Width - m.config.PaddingLeft - m.config.PaddingRight
	trackX := m.config.PaddingLeft + (availableWidth-m.config.TrackConfig.Width)/2

	shadowOffset := m.config.TrackConfig.ShadowOffset

	w = m.config.TrackConfig.Width + 2*shadowOffset
	h = trackHeight + 2*shadowOffset
	x = trackX - shadowOffset
	y = m.config.PaddingTop - shadowOffset

	return
}

func (m *Theme) GetZeroDBLinePosition(
	displayRatio float32,
	size fyne.Size,
) (xStart, yStart, xEnd, yEnd float32) {
	trackHeight := size.Height - m.config.PaddingTop - m.config.PaddingBottom
	zeroY := m.config.PaddingTop + (1.0-displayRatio)*trackHeight

	availableWidth := size.Width - m.config.PaddingLeft - m.config.PaddingRight
	trackX := m.config.PaddingLeft + (availableWidth-m.config.TrackConfig.Width)/2

	xStart = trackX - 5
	yStart = zeroY
	xEnd = trackX + m.config.TrackConfig.Width + 5
	yEnd = zeroY

	return
}

func (m *Theme) GetLabelXPosition() float32 {
	labelWidth := m.config.TrackConfig.LabelWidth
	paddingLeft := m.config.PaddingLeft

	labelX := paddingLeft - labelWidth - 6
	if labelX < 0 {
		labelX = 0
	}

	return labelX
}
