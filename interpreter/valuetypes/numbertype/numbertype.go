package numbertype

import "fmt"

type NumberType struct {
	value float32
}

func New(value float32) NumberType {
	return NumberType{value: value}
}

func (nt NumberType) Fmt() string {
	return fmt.Sprint(nt.value)
}

func (nt NumberType) Lit() any {
	return nt.value
}
