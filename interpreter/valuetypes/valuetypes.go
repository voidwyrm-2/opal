package valuetypes

type ValueType interface {
	Fmt() string
	Lit() any
	Type() string
	Add(val ValueType) (ValueType, error)
	Concat(val ValueType) (ValueType, error)
	Sub(val ValueType) (ValueType, error)
	Mul(val ValueType) (ValueType, error)
	Div(val ValueType) (ValueType, error)
	Mod(val ValueType) (ValueType, error)
	BitAnd(val ValueType) (ValueType, error)
	BitOr(val ValueType) (ValueType, error)
	BitXOR(val ValueType) (ValueType, error)
	Equals(val ValueType) (ValueType, error)
	NotEquals(val ValueType) (ValueType, error)
	GreaterThan(val ValueType) (ValueType, error)
	LesserThan(val ValueType) (ValueType, error)
}
