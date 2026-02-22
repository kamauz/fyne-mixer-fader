# Technical Documentation

## API Reference

### Constructor

```go
func mixerfader.New(
    value      float64,
    tapeType   enums.TapeTypeEnum,
    themeConfig config.ThemeConfig,
    onChanged  func(float64),
) *MixerFader
```

| Parameter     | Type                 | Description                                                |
| ------------- | -------------------- | ---------------------------------------------------------- |
| `value`       | `float64`            | Initial value in dB                                        |
| `tapeType`    | `enums.TapeTypeEnum` | Response curve (`TapeTypeLinear` or `TapeTypeLogarithmic`) |
| `themeConfig` | `config.ThemeConfig` | Visual styling                                             |
| `onChanged`   | `func(float64)`      | Called on every value change                               |

### MixerFader

```go
type MixerFader struct {
    Value     float64        // current value in dB (read/write)
    OnChanged func(float64)  // value change callback
}
```

### Tape types

| Constant                    | Behaviour                                                                    |
| --------------------------- | ---------------------------------------------------------------------------- |
| `enums.TapeTypeLinear`      | Uniform dB distribution across the full range (−120 to +12)                  |
| `enums.TapeTypeLogarithmic` | Non-linear curve with finer resolution at lower dB — typical mixer behaviour |

The logarithmic tape divides the slider into four equal segments that map to unequal dB ranges:

| Segment (position) | dB range       |
| ------------------ | -------------- |
| 0–25%              | −120 to −60 dB |
| 25–50%             | −60 to −20 dB  |
| 50–75%             | −20 to 0 dB    |
| 75–100%            | 0 to +12 dB    |

---

## Theme Configuration

Each fader accepts its own `config.ThemeConfig`, so multiple faders in the same layout can have different styles.

```go
type ThemeConfig struct {
    TrackConfig   TrackStyle
    ThumbConfig   ThumbStyle
    PaddingTop    float32
    PaddingRight  float32
    PaddingBottom float32
    PaddingLeft   float32
}
```

### TrackStyle

```go
type TrackStyle struct {
    Color           color.NRGBA  // track fill color
    ZeroDBLineColor color.NRGBA  // horizontal marker at 0 dB
    SideScaleColor  color.NRGBA  // dB scale label color
    ShadowColor     color.NRGBA  // track shadow color
    ShadowOffset    float32      // shadow spread in pixels
    Width           float32      // track width in pixels
    Radius          float32      // track corner radius
    LabelWidth      float32      // width of each scale label cell
    LabelHeight     float32      // height of each scale label cell
}
```

### ThumbStyle

```go
type ThumbStyle struct {
    Color           color.NRGBA  // thumb fill color
    BorderColor     color.NRGBA  // thumb border color
    MiddleLineColor color.NRGBA  // center line drawn across the thumb
    LabelColor      color.NRGBA  // value label text color
    LabelBoxColor   color.NRGBA  // value label background color
    ShowLabel       bool         // show value label on hover
    ShowLabelBox    bool         // show background box behind the label
    Width           float32      // thumb width in pixels
    Height          float32      // thumb height in pixels
}
```

---

## Architecture

```
mixerfader/
  slider.go           — MixerFader widget, constructor, CreateRenderer
  events.go           — Tapped, Dragged, DoubleTapped, MouseIn/Out
  tape.go             — Thin wrapper methods delegating to the active tape
  renderer.go         — Fyne WidgetRenderer implementation (Layout, Refresh, Objects)
  renderer_helper.go  — Per-element render functions (track, thumb, labels, scale)

  enums/
    tapetype.go       — TapeTypeEnum constants (Linear, Logarithmic)

  tapes/
    factory.go        — CreateTapeFactory: selects and instantiates the right tape
    models/
      scale.go        — ScaleLabelSpec: position and text for one scale label
    types/
      interface.go    — Tapable interface
      linear.go       — linearTape implementation
      logaritmic.go   — logarithmicTape implementation

  theme/
    theme.go          — Theme struct, NewTheme constructor
    getters.go        — Geometry calculations (positions, sizes)
    setters.go        — Style application (colors, corner radius, stroke)
    config/
      theme.go        — ThemeConfig struct
      track.go        — TrackStyle struct
      thumb.go        — ThumbStyle struct
```

### Key design decisions

**Tape / Theme separation.** A tape is pure math — it converts dB values to display ratios and Y positions to values. It has no knowledge of the Theme or any canvas element. Geometry parameters (trackHeight, paddingTop, label dimensions) are passed in by the renderer at call time.

**Renderer owns canvas objects.** All `canvas.*` objects are created once in `CreateRenderer` and updated in-place on every `Layout` call. No objects are created during `Refresh`.

**Per-widget theme.** `ThemeConfig` is passed by value to `New()`, so each fader instance owns its own independent copy. Changing one fader's style does not affect others.

---

## Implementing a custom tape

Implement the `Tapable` interface and register it in `factory.go`:

```go
type Tapable interface {
    GetMin() float32
    GetMax() float32
    GetDefaultValue() float32
    FromYPositionToValue(yPos, trackHeight, paddingTop float32) float32
    FromValueToDisplayRatio(value float32) float32
    CalculateSideScale(trackHeight, paddingTop, labelWidth, labelHeight, labelX float32) []models.ScaleLabelSpec
}
```

`FromValueToDisplayRatio` and `FromYPositionToValue` must be strict inverses of each other: for any valid dB value `v`, the round-trip `FromYPositionToValue(paddingTop + (1−FromValueToDisplayRatio(v))×trackHeight, trackHeight, paddingTop)` must return `v` within floating-point tolerance.

---

## Running tests

```sh
go test ./mixerfader/tapes/...
```

Tests cover both tape implementations (getters, ratio boundaries, Y↔value round-trips, scale label positions and text) and the factory (known types, panic on unknown type).
