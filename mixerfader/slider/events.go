package mixerfader

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// Tapped implements fyne.Tappable
func (s *MixerFader) Tapped(ev *fyne.PointEvent) { s.applyY(ev.Position.Y) }
func (s *MixerFader) DoubleTapped(ev *fyne.PointEvent) {
	newValue := float64(s.tape.GetDefaultValue())
	s.SetDB(newValue)
}

// Dragged implements fyne.Draggable
func (s *MixerFader) Dragged(ev *fyne.DragEvent) { s.applyY(ev.Position.Y) }

// DragEnd implements fyne.Draggable
func (s *MixerFader) DragEnd() {}

// MouseIn implements desktop.Hoverable
func (s *MixerFader) MouseIn(*desktop.MouseEvent) {
	s.hovered = true
	s.Refresh()
}

// MouseMoved implements desktop.Hoverable
func (s *MixerFader) MouseMoved(*desktop.MouseEvent) {}

// MouseOut implements desktop.Hoverable
func (s *MixerFader) MouseOut() {
	s.hovered = false
	s.Refresh()
}

///// utils ///////

func (s *MixerFader) applyY(y float32) {
	// get paddings
	paddingTop, _, _, _ := s.theme.GetPaddings()

	// get new value after drag
	trackHeight := s.theme.GetTrackHeight(s.Size().Height)
	newValue := s.tape.FromYPositionToValue(y, trackHeight, paddingTop)

	// set new value
	s.SetDB(float64(newValue))
}
