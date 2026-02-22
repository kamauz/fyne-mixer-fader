package mixerfader

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/models"
)

func (r *renderer) renderTrack(
	size fyne.Size,
) {
	// set track style
	r.theme.SetTrack(r.track)

	// get track position
	w, h, x, y := r.theme.GetTrackPosition(size)

	// render
	r.track.Resize(fyne.NewSize(w, h))
	r.track.Move(fyne.NewPos(x, y))
}

// SetTrackShadow positions and styles the shadow effect behind the track element.
func (r *renderer) renderTrackShadow(
	size fyne.Size,
) {
	// set track shadow style
	r.theme.SetTrackShadow(r.trackShadow)

	// get track shadow info
	w, h, x, y := r.theme.GetTrackShadow(size)

	// render
	r.trackShadow.Resize(fyne.NewSize(w, h))
	r.trackShadow.Move(fyne.NewPos(x, y))
}

func (r *renderer) renderSideScale(
	size fyne.Size,
) {

	trackHeight := r.theme.GetTrackHeight(size.Height)

	// get track label size
	labelWidth, labelHeight := r.theme.GetTrackLabelSize()

	// get label x position
	labelX := r.theme.GetLabelXPosition()

	// get paddings
	paddingTop, _, _, _ := r.theme.GetPaddings()

	// calculate side scale
	var scaleValues []models.ScaleLabelSpec = r.slider.CalculateTapeSideScale(
		trackHeight,
		paddingTop,
		labelWidth,
		labelHeight,
		labelX,
	)

	// render
	for i, v := range scaleValues {
		if len(r.scaleLabels) <= i {
			// Create new label if not enough
			lbl := canvas.NewText("", color.Transparent)
			r.theme.SetTrackSideScaleLabel(lbl)
			lbl.Text = v.Text

			// render
			lbl.Move(fyne.NewPos(v.X, v.Y))
			lbl.Resize(fyne.NewSize(v.Width, v.Height))
			lbl.Show()
			r.scaleLabels = append(r.scaleLabels, lbl)
		} else {
			lbl := r.scaleLabels[i]
			r.theme.SetTrackSideScaleLabel(lbl)
			lbl.Text = v.Text

			// render
			lbl.Move(fyne.NewPos(v.X, v.Y))
			lbl.Resize(fyne.NewSize(v.Width, v.Height))
			lbl.Show()
		}
	}

	// hide extra values
	for i := len(scaleValues); i < len(r.scaleLabels); i++ {
		r.scaleLabels[i].Hide()
	}
}

func (r *renderer) renderZeroDBLine(
	size fyne.Size,
) {
	// get 0db display ratio position w.r.t. the tape
	displayRatio := r.slider.FromValueToTapeDisplayRatio(0) // 0 dB

	// get zero db line position
	xStart, yStart, xEnd, yEnd := r.theme.GetZeroDBLinePosition(
		displayRatio,
		size,
	)

	// render
	r.zeroDBLine.Position1 = fyne.NewPos(xStart, yStart)
	r.zeroDBLine.Position2 = fyne.NewPos(xEnd, yEnd)
}
