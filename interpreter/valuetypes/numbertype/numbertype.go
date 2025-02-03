package numbertype

import (
	"fmt"

	"github.com/voidwyrm-2/opal/interpreter/valuetypes"
)

type NumberType struct {
	value float32
}

func New(value float32) NumberType {
	return NumberType{value: value}
}

func FromBool(value bool) NumberType {
	if value {
		return NumberType{value: 1}
	}
	return NumberType{value: 0}
}

func (nt NumberType) Fmt() string {
	return fmt.Sprint(nt.value)
}

func (nt NumberType) Lit() any {
	return nt.value
}

func (nt NumberType) Type() string {
	return "number"
}

func (nt NumberType) Add(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) Concat(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) Sub(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) Mul(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) Div(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) Mod(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) BitAnd(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) BitOr(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) BitXOR(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) Equals(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, nil
}

func (nt NumberType) NotEquals(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	return nil, nil
}

func (nt NumberType) GreaterThan(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	b := val.(NumberType)
	return FromBool(nt.value > b.value), nil
}

func (nt NumberType) LesserThan(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() != "number" {
		return val.Add(nt)
	}
	b := val.(NumberType)
	return FromBool(nt.value < b.value), nil
}
