package conversions

import (
	"fmt"
)

type Unit string

type Amount struct {
	Value float64
	Unit  Unit
}

func (a Amount) String() string {
	modifier := ""
	modifiedValue := a.Value
	return fmt.Sprintf("%f %s%s", modifiedValue, modifier, a.Unit)
}

type Conversion struct {
	Input  Amount
	Output Amount
}

type Converter func(Amount) Conversion
