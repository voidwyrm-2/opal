package nodes

import "github.com/voidwyrm-2/opal/interpreter/valuetypes"

type Node interface {
	Execute(vars map[string]valuetypes.ValueType) error
	Str() string
}
