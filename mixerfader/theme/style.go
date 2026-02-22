package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (m *Theme) SetTrack(track *canvas.Rectangle) {
	// get theme track config
	themeConfig := m.config.TrackConfig

	// set style
	track.CornerRadius = themeConfig.Radius
	track.FillColor = themeConfig.Color
}

func (m *Theme) SetTrackSideScaleLabel(label *canvas.Text) {
	// get theme track config
	themeConfig := m.config.TrackConfig.SideScaleStyle

	// set style
	label.Color = themeConfig.Color
	label.TextSize = themeConfig.TextSize
}

// SetTrackShadow positions and styles the shadow effect behind the track element.
func (m *Theme) SetTrackShadow(trackShadow *canvas.Rectangle) {
	// get theme track config
	themeConfig := m.config.TrackConfig.ShadowStyle
	trackConfig := m.config.TrackConfig

	// set style
	trackShadow.FillColor = themeConfig.Color
	trackShadow.CornerRadius = trackConfig.Radius + themeConfig.Offset
}

func (m *Theme) SetThumb(thumb *canvas.Rectangle) {
	// get theme thumb config
	themeConfig := m.config.ThumbConfig

	// set stule
	thumb.CornerRadius = themeConfig.CornerRadius
	thumb.FillColor = themeConfig.Color
	thumb.StrokeColor = themeConfig.BorderColor
	thumb.StrokeWidth = themeConfig.StrokeWidth

	// set size
	thumbWidth := m.config.ThumbConfig.Width
	thumbHeight := m.config.ThumbConfig.Height
	thumb.Resize(fyne.NewSize(thumbWidth, thumbHeight))
}

// SetThumbLabel positions and styles the value label text displayed above the thumb.
func (m *Theme) SetThumbLabel(valueLabel *canvas.Text) {
	// get theme thumb label style
	themeConfig := m.config.ThumbConfig.LabelStyle

	// set style
	valueLabel.Color = themeConfig.Color
	valueLabel.Alignment = themeConfig.TextAlign
	valueLabel.TextStyle.Bold = themeConfig.Bold
}

// SetThumbLabelBox positions and styles the background box container for the value label.
func (m *Theme) SetThumbLabelBox(valueLabelContainer *canvas.Rectangle) {
	// get theme thumb middle line style
	themeConfig := m.config.ThumbConfig.LabelBoxStyle

	// set style
	valueLabelContainer.FillColor = themeConfig.Color
	valueLabelContainer.CornerRadius = themeConfig.CornerRadius
}

func (m *Theme) SetMiddleThumbLine(middleThumbLine *canvas.Line) {
	if middleThumbLine == nil {
		return
	}
	// get theme thumb middle line style
	themeConfig := m.config.ThumbConfig.MiddleLineStyle

	// set style
	middleThumbLine.StrokeColor = themeConfig.StrokeColor
	middleThumbLine.StrokeWidth = themeConfig.StrokeWidth
}
