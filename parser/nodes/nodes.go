package nodes

import "github.com/voidwyrm-2/opal/interpreter/scope"

type Node interface {
	Execute(vars map[string]string) error
	Str() string
}
