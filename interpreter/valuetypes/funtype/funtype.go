package funtype

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/voidwyrm-2/opal/common"
	"github.com/voidwyrm-2/opal/interpreter/valuetypes"
	"github.com/voidwyrm-2/opal/interpreter/valuetypes/numbertype"
)

type FunType struct {
	value func(args []valuetypes.ValueType)
	hash  string // used for equality reasons
}

func New(value func(args []valuetypes.ValueType)) FunType {
	return FunType{value: value, hash: fmt.Sprintf("%d%d%d%d%d", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))}
}

func (ft FunType) Fmt() string {
	return "fun<" + ft.hash + ">"
}

func (ft FunType) Lit() any {
	return ft.value
}

func (ft FunType) Type() string {
	return "fun"
}

func (ft FunType) Add(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support addition")
}

func (ft FunType) Concat(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, nil
}

func (ft FunType) Sub(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support subtraction")
}

func (ft FunType) Mul(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support multiplication")
}

func (ft FunType) Div(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support division")
}

func (ft FunType) Mod(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support modulus")
}

func (ft FunType) BitAnd(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support bitwise AND")
}

func (ft FunType) BitOr(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support bitwise OR")
}

func (ft FunType) BitXOR(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support bitwise XOR")
}

func (ft FunType) Equals(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() == "fun" {
		fn, _ := val.(FunType)
		if ft.hash == fn.hash {
			return numbertype.New(1), nil
		}
	}

	return numbertype.New(0), nil
}

func (ft FunType) NotEquals(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return func() valuetypes.ValueType {
		if common.Assert(ft.Equals(val)).Lit() == 1 {
			return numbertype.New(0)
		} else {
			return numbertype.New(1)
		}
	}(), nil
}

func (ft FunType) GreaterThan(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support greater than")
}

func (ft FunType) LesserThan(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + ft.Type() + "' does not support lesser than")
}
