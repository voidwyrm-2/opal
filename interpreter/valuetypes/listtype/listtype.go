package listtype

import (
	"errors"
	"strings"

	"github.com/voidwyrm-2/opal/interpreter/valuetypes"
)

type Node struct {
	value      valuetypes.ValueType
	next, prev *Node
}

func NewNode(value valuetypes.ValueType, next, prev *Node) *Node {
	return &Node{value: value, next: next, prev: prev}
}

type ListType struct {
	back, front *Node
	length      int
}

func New(initialValues ...valuetypes.ValueType) ListType {
	l := ListType{back: nil, front: nil, length: 0}
	for _, v := range initialValues {
		l.Append(v)
	}
	return l
}

func (lt *ListType) Iter(fn func(val valuetypes.ValueType) error) error {
	if lt.length == 0 {
		return nil
	}

	current := lt.back
	for current != nil {
		if err := fn(current.value); err != nil {
			return err
		}
		current = current.next
	}

	return nil
}

func (lt *ListType) Map(fn func(val valuetypes.ValueType) (valuetypes.ValueType, error)) (ListType, error) {
	mapped := New()
	if lt.length == 0 {
		return mapped, nil
	}

	current := lt.back
	for current != nil {
		if v, err := fn(current.value); err != nil {
			return ListType{}, err
		} else {
			mapped.Append(v)
		}
		current = current.next
	}

	return mapped, nil
}

func (lt *ListType) Append(value valuetypes.ValueType) {
	if lt.length == 0 {
		lt.front = NewNode(value, nil, nil)
		lt.back = lt.front
	} else {
		oldFront := lt.front
		newFront := NewNode(value, nil, nil)
		newFront.prev = oldFront
		oldFront.next = newFront
		lt.front = newFront
	}
	lt.length++
}

func (lt *ListType) Pop() (valuetypes.ValueType, error) {
	if lt.length == 0 {
		return nil, errors.New("cannot empty pop from an empty list")
	}
	lt.length--
	popped := lt.front
	lt.front, popped.prev = popped.prev, nil
	if lt.length == 0 {
		lt.back = nil
	}
	return nil, nil
}

func (lt ListType) Fmt() string {
	formatted := []string{}
	lt.Iter(func(val valuetypes.ValueType) error {
		formatted = append(formatted, val.Fmt())
		return nil
	})
	return "[ " + strings.Join(formatted, ", ") + " ]"
}

func (lt ListType) Lit() any {
	return lt
}

func (lt ListType) Type() string {
	return "list"
}

func (lt ListType) Add(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.Add(val)
	})
}

func (lt ListType) Concat(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	if val.Type() == "List" {
		l, _ := val.(ListType)
		l.Iter(func(val valuetypes.ValueType) error {
			lt.Append(val)
			return nil
		})
	} else {
		lt.Append(val)
	}
	return lt, nil
}

func (lt ListType) Sub(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.Sub(val)
	})
}

func (lt ListType) Mul(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.Mul(val)
	})
}

func (lt ListType) Div(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.Div(val)
	})
}

func (lt ListType) Mod(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.Mod(val)
	})
}

func (lt ListType) BitAnd(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.BitAnd(val)
	})
}

func (lt ListType) BitOr(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.BitOr(val)
	})
}

func (lt ListType) BitXOR(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return lt.Map(func(subval valuetypes.ValueType) (valuetypes.ValueType, error) {
		return subval.BitXOR(val)
	})
}

func (lt ListType) Equals(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	panic("unimplemented")
	return nil, nil
}

func (lt ListType) NotEquals(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	panic("unimplemented")
	return nil, nil
}

func (lt ListType) GreaterThan(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + lt.Type() + "' does not support greater than")
}

func (lt ListType) LesserThan(val valuetypes.ValueType) (valuetypes.ValueType, error) {
	return nil, errors.New("type '" + lt.Type() + "' does not support lesser than")
}
