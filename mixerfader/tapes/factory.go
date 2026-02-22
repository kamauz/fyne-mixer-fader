package tapes

import (
	"fmt"

	"github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/enums"
	tapeTypes "github.com/kamauz/fyne-mixer-fader/mixerfader/tapes/types"
)

func CreateTapeFactory(
	tapeType enums.TapeTypeEnum,
) (tapeTypes.Tapable, error) {

	var defaultTape tapeTypes.Tapable
	switch tapeType {

	// most typical case
	case enums.TapeTypeLogarithmic:
		defaultTape = tapeTypes.NewLogarithmicTape()
	case enums.TapeTypeLinear:
		defaultTape = tapeTypes.NewLinearTape()
	default:
		return nil, fmt.Errorf("tape type not supported: %q", tapeType)
	}

	return defaultTape, nil
}
