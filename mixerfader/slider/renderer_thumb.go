package mixerfader

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (r *renderer) renderThumb(
	size fyne.Size,
	thumbYAnchor *float32,
) {
	// set thumb style
	r.theme.SetThumb(r.thumb)

	// get ratio from tape
	displayRatio := r.slider.FromValueToTapeDisplayRatio(float32(r.slider.DB()))

	// get thumb position
	thumbX := r.theme.GetThumbPosition(
		displayRatio,
		thumbYAnchor,
		r.thumb.Size(),
		size,
	)

	// move thumb
	r.thumb.Move(fyne.NewPos(thumbX, *thumbYAnchor))
}

func (r *renderer) renderMiddleThumbLine(
	posY *float32,
) {
	// set style
	r.theme.SetMiddleThumbLine(r.middleThumbLine)

	// get new middle thumb line position
	lineStartX, lineEndX, lineY := r.theme.GetThumbMiddleLinePosition(
		*posY,
		r.thumb.Size(),
	)

	// move thumb middle line
	r.middleThumbLine.Position1 = fyne.NewPos(lineStartX, lineY)
	r.middleThumbLine.Position2 = fyne.NewPos(lineEndX, lineY)
}

// SetThumbLabel positions and styles the value label text displayed above the thumb.
func (r *renderer) renderThumbLabel(
	size fyne.Size,
	thumbYAnchor *float32,
) {
	// set value label style
	r.theme.SetThumbLabel(r.valueLabel)

	// retrieve thumb label position
	labelX, labelWidth := r.theme.GetThumbLabelPosition(size)

	value := r.slider.DB()
	min := float64(r.slider.GetTapeMinValue())

	if value == min {
		r.valueLabel.Text = "-inf"
	} else {
		r.valueLabel.Text = fmt.Sprintf("%.1f db", value)
	}

	// set position to thumb label in according to the thumb position
	r.valueLabel.Move(fyne.NewPos(labelX, *thumbYAnchor-28))
	r.valueLabel.Resize(fyne.NewSize(labelWidth, 20))
}

// SetThumbLabelBox positions and styles the background box container for the value label.
func (r *renderer) renderThumbLabelBox(
	size fyne.Size,
	thumbYAnchor *float32,
) {
	// set thumb label box style
	r.theme.SetThumbLabelBox(r.valueLabelContainer)

	// get thumb label box position
	w, h, x, y := r.theme.GetThumbLabelBoxPosition(
		thumbYAnchor,
		size,
	)

	// set position in according to the thumb position
	r.valueLabelContainer.Move(fyne.NewPos(x, y))
	r.valueLabelContainer.Resize(fyne.NewSize(w, h))
}
