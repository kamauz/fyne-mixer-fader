package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	fyneWidget "fyne.io/fyne/v2/widget"

	mixerfader "github.com/kamauz/fyne-mixer-fader/mixerfader/slider"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/enums"
	"github.com/kamauz/fyne-mixer-fader/mixerfader/theme/config"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	themeConfig := config.ThemeConfig{
		TrackConfig: config.TrackStyle{
			Color:           color.NRGBA{R: 22, G: 22, B: 22, A: 255},
			ZeroDBLineColor: color.NRGBA{R: 255, G: 200, B: 0, A: 180},
			ShadowColor:     color.NRGBA{R: 200, G: 200, B: 200, A: 140},
			ShadowOffset:    float32(3),
			Width:           float32(7),
			Radius:          float32(3.5),
			LabelWidth:      float32(50),
			LabelHeight:     float32(18),

			// children
			SideScaleStyle: config.SideScaleStyle{
				Color:    color.NRGBA{R: 250, G: 250, B: 250, A: 255},
				TextSize: 10,
			},
			ShadowStyle: config.ShadowStyle{
				Color:  color.NRGBA{R: 200, G: 200, B: 200, A: 140},
				Offset: float32(3),
			},
		},
		ThumbConfig: config.ThumbStyle{
			ShowLabel:    true,
			ShowLabelBox: true,
			Color:        color.NRGBA{R: 255, G: 255, B: 255, A: 220},
			BorderColor:  color.NRGBA{R: 200, G: 200, B: 200, A: 190},
			Width:        float32(32),
			Height:       float32(46),
			CornerRadius: float32(4),
			StrokeWidth:  float32(2),

			// children
			MiddleLineStyle: config.MiddleLineStyle{
				StrokeColor: color.NRGBA{R: 0, G: 0, B: 0, A: 255},
				StrokeWidth: 2,
			},
			LabelBoxStyle: config.LabelBoxStyle{
				Color:        color.NRGBA{R: 22, G: 180, B: 22, A: 220},
				CornerRadius: float32(6),
			},
			LabelStyle: config.LabelStyle{
				Color:     color.NRGBA{R: 250, G: 250, B: 250, A: 240},
				TextAlign: fyne.TextAlignCenter,
				Bold:      true,
			},
		},
		PaddingTop:    float32(10),
		PaddingRight:  float32(0),
		PaddingBottom: float32(10),
		PaddingLeft:   float32(30),
	}

	// Creazione degli slider personalizzati
	sliders := []fyne.CanvasObject{
		container.NewPadded(mixerfader.New(0, enums.TapeTypeLinear, themeConfig, func(a float64) {
			fmt.Println("new value", a)
		})),
		container.NewPadded(mixerfader.New(0, enums.TapeTypeLogarithmic, themeConfig, func(a float64) {})),
		container.NewPadded(mixerfader.New(-20, enums.TapeTypeLinear, themeConfig, func(a float64) {})),
		container.NewPadded(mixerfader.New(0, enums.TapeTypeLogarithmic, themeConfig, func(a float64) {})),
		container.NewPadded(mixerfader.New(-45.3, enums.TapeTypeLogarithmic, themeConfig, func(a float64) {})),
	}

	w.SetContent(container.NewVBox(
		fyneWidget.NewButton("Click me", func() {}),
		container.NewHBox(sliders...),
	))
	w.ShowAndRun()
}
