package test

import (
	"testing"

	tapeType "github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/types"
	"github.com/stretchr/testify/assert"
)

// Costanti geometriche usate in tutti i test della linear tape.
// Simulano una track alta 300px con 10px di padding in cima.
const (
	linTrackH = float32(300)
	linPadTop = float32(10)
)

// ─── Getters ──────────────────────────────────────────────────────────────────

func TestLinear_GetMin(t *testing.T) {
	assert.Equal(t, float32(-120), tapeType.NewLinearTape().GetMin())
}

func TestLinear_GetMax(t *testing.T) {
	assert.Equal(t, float32(12), tapeType.NewLinearTape().GetMax())
}

func TestLinear_GetDefaultValue(t *testing.T) {
	assert.Equal(t, float32(0), tapeType.NewLinearTape().GetDefaultValue())
}

// ─── FromValueToDisplayRatio ──────────────────────────────────────────────────

func TestLinear_Ratio_MinGivesZero(t *testing.T) {
	tape := tapeType.NewLinearTape()
	assert.InDelta(t, 0.0, tape.FromValueToDisplayRatio(tape.GetMin()), 1e-5)
}

func TestLinear_Ratio_MaxGivesOne(t *testing.T) {
	tape := tapeType.NewLinearTape()
	assert.InDelta(t, 1.0, tape.FromValueToDisplayRatio(tape.GetMax()), 1e-5)
}

func TestLinear_Ratio_MidpointIsHalf(t *testing.T) {
	tape := tapeType.NewLinearTape()
	mid := (tape.GetMin() + tape.GetMax()) / 2
	assert.InDelta(t, 0.5, tape.FromValueToDisplayRatio(mid), 1e-4)
}

func TestLinear_Ratio_ZeroDB(t *testing.T) {
	tape := tapeType.NewLinearTape()
	expected := float32(120.0 / 132.0)
	assert.InDelta(t, expected, tape.FromValueToDisplayRatio(0), 1e-4)
}

func TestLinear_Ratio_ClampBelowMin(t *testing.T) {
	assert.Equal(t, float32(0), tapeType.NewLinearTape().FromValueToDisplayRatio(-999))
}

func TestLinear_Ratio_ClampAboveMax(t *testing.T) {
	assert.Equal(t, float32(1), tapeType.NewLinearTape().FromValueToDisplayRatio(999))
}

// ─── FromYPositionToValue ─────────────────────────────────────────────────────

func TestLinear_Y_TopOfTrackGivesMax(t *testing.T) {
	tape := tapeType.NewLinearTape()
	val := tape.FromYPositionToValue(linPadTop, linTrackH, linPadTop)
	assert.InDelta(t, tape.GetMax(), val, 0.01)
}

func TestLinear_Y_BottomOfTrackGivesMin(t *testing.T) {
	tape := tapeType.NewLinearTape()
	val := tape.FromYPositionToValue(linPadTop+linTrackH, linTrackH, linPadTop)
	assert.InDelta(t, tape.GetMin(), val, 0.01)
}

func TestLinear_Y_MiddleGivesMidValue(t *testing.T) {
	tape := tapeType.NewLinearTape()
	midY := linPadTop + linTrackH/2
	val := tape.FromYPositionToValue(midY, linTrackH, linPadTop)
	expected := tape.GetMin() + (tape.GetMax()-tape.GetMin())/2
	assert.InDelta(t, expected, val, 0.5)
}

func TestLinear_Y_AboveTrackClampsToMax(t *testing.T) {
	tape := tapeType.NewLinearTape()
	// Y = 0, che è prima del paddingTop
	val := tape.FromYPositionToValue(0, linTrackH, linPadTop)
	assert.Equal(t, tape.GetMax(), val)
}

func TestLinear_Y_BelowTrackClampsToMin(t *testing.T) {
	tape := tapeType.NewLinearTape()
	val := tape.FromYPositionToValue(linPadTop+linTrackH+100, linTrackH, linPadTop)
	assert.Equal(t, tape.GetMin(), val)
}

func TestLinear_Y_ZeroTrackHeightGuard(t *testing.T) {
	tape := tapeType.NewLinearTape()
	// Con trackHeight = 0 non deve andare in divisione per zero
	val := tape.FromYPositionToValue(50, 0, linPadTop)
	assert.Equal(t, tape.GetMin(), val)
}

// ─── Round-trip: value → ratio → Y → value ───────────────────────────────────

func TestLinear_RoundTrip(t *testing.T) {
	tape := tapeType.NewLinearTape()
	cases := []struct {
		name  string
		value float32
	}{
		{"min", -120},
		{"low", -84},
		{"middle", -54},
		{"zero_dB", 0},
		{"max", 12},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ratio := tape.FromValueToDisplayRatio(tc.value)
			yPos := linPadTop + (1.0-ratio)*linTrackH
			got := tape.FromYPositionToValue(yPos, linTrackH, linPadTop)
			assert.InDelta(t, tc.value, got, 0.5, "round-trip failed for %.1f dB", tc.value)
		})
	}
}

// ─── CalculateSideScale ───────────────────────────────────────────────────────

func TestLinear_Scale_HasZeroDBLabel(t *testing.T) {
	tape := tapeType.NewLinearTape()
	specs := tape.CalculateSideScale(linTrackH, linPadTop, 50, 18, 0)
	found := false
	for _, s := range specs {
		if s.Text == "0 dB" {
			found = true
			break
		}
	}
	assert.True(t, found, "there must be a label '0 dB'")
}

func TestLinear_Scale_AllLabelsHaveCorrectX(t *testing.T) {
	tape := tapeType.NewLinearTape()
	labelX := float32(7)
	specs := tape.CalculateSideScale(linTrackH, linPadTop, 50, 18, labelX)
	for _, s := range specs {
		assert.Equal(t, labelX, s.X)
	}
}

func TestLinear_Scale_AllLabelsHaveCorrectDimensions(t *testing.T) {
	tape := tapeType.NewLinearTape()
	labelW, labelH := float32(50), float32(18)
	specs := tape.CalculateSideScale(linTrackH, linPadTop, labelW, labelH, 0)
	for _, s := range specs {
		assert.Equal(t, labelW, s.Width)
		assert.Equal(t, labelH, s.Height)
	}
}
