package enums

type TapeTypeEnum string

const (
	TapeTypeLogarithmic TapeTypeEnum = "LOGARITHMIC"
	TapeTypeLinear      TapeTypeEnum = "LINEAR"
)

func (t TapeTypeEnum) String() string {
	return string(t)
}

func (t TapeTypeEnum) IsValid() bool {
	switch t {
	case TapeTypeLogarithmic, TapeTypeLinear:
		return true
	default:
		return false
	}
}
