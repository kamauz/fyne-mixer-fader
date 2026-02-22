package mixerfader

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	fadertheme "github.com/kamauz/fyne-mixer-fader/mixerfader/theme"
)

// customize renderer
type renderer struct {
	slider *MixerFader

	// customization
	background          *canvas.Rectangle
	track               *canvas.Rectangle
	thumb               *canvas.Rectangle
	trackShadow         *canvas.Rectangle
	valueLabelContainer *canvas.Rectangle
	valueLabel          *canvas.Text

	theme   *fadertheme.Theme
	objects []fyne.CanvasObject

	// dB scale
	scaleLabels []*canvas.Text
	zeroDBLine  *canvas.Line

	middleThumbLine *canvas.Line
}

func (r *renderer) MinSize() fyne.Size {
	// get theme min size
	w, h := r.theme.GetMinSize()
	return fyne.NewSize(w, h)
}

func (r *renderer) Layout(size fyne.Size) {
	// layout background
	r.background.Resize(size)
	r.background.Move(fyne.NewPos(0, 0))

	// define track (offset by paddingLeft)
	r.renderTrack(size)

	// define track shadow
	r.renderTrackShadow(size)

	// slider movable y position
	var posY float32

	// define thumb
	r.renderThumb(
		size, &posY,
	)

	r.renderMiddleThumbLine(&posY)

	if r.theme.IsShowLabelBox() {
		r.renderThumbLabelBox(size, &posY)
		r.valueLabelContainer.Hidden = !r.slider.hovered
	}

	if r.theme.IsShowLabel() {
		r.renderThumbLabel(size, &posY)
		r.valueLabel.Hidden = !r.slider.hovered
	}

	// render side scale
	r.renderSideScale(size)

	// rendere 0db line
	r.renderZeroDBLine(size)
}

func (r *renderer) Refresh() {
	r.Layout(r.slider.Size())
	canvas.Refresh(r.slider)
}

func (r *renderer) Objects() []fyne.CanvasObject {
	objects := []fyne.CanvasObject{
		r.background,
		r.trackShadow,
		r.track,
	}

	// add 0dB line before thumb so it appears behind
	objects = append(objects, r.zeroDBLine)

	// add thumb
	objects = append(objects,
		r.thumb,
	)

	// Add black horizontal line in the middle of the thumb
	if r.middleThumbLine != nil {
		objects = append(objects, r.middleThumbLine)
	}

	objects = append(objects,
		r.valueLabelContainer,
		r.valueLabel,
	)

	// add scale labels and markers
	for _, label := range r.scaleLabels {
		objects = append(objects, label)
	}

	return objects
}

func (r *renderer) Destroy() {}

// creator
func newMixerFaderRenderer(
	component *MixerFader,

	// theme
	theme *fadertheme.Theme,
) *renderer {

	// Declare canvas elements
	background := canvas.NewRectangle(color.Transparent)
	track := canvas.NewRectangle(color.Transparent)
	trackShadow := canvas.NewRectangle(color.Transparent)
	thumb := canvas.NewRectangle(color.Transparent)
	valueLabel := canvas.NewText("", color.Transparent)
	valueLabelContainer := canvas.NewRectangle(color.Transparent)

	// Create enough labels/markers for maximum range (e.g., 30 markers)
	const maxScaleLabels = 30
	scaleLabels := make([]*canvas.Text, maxScaleLabels)
	for i := range scaleLabels {
		scaleLabels[i] = canvas.NewText("", color.Transparent)
	}

	zeroDBLine := canvas.NewLine(color.Transparent)
	middleThumbLine := canvas.NewLine(color.Transparent)

	return &renderer{
		// fader
		slider: component,

		// configs
		theme: theme,

		objects: []fyne.CanvasObject{},

		// ui components
		background:          background,
		track:               track,
		thumb:               thumb,
		trackShadow:         trackShadow,
		valueLabelContainer: valueLabelContainer,
		valueLabel:          valueLabel,
		scaleLabels:         scaleLabels,
		zeroDBLine:          zeroDBLine,
		middleThumbLine:     middleThumbLine,
	}
}
