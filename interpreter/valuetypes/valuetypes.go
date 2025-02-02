package valuetypes

type ValueType interface {
	Fmt() string
	Lit() any
	Add(st ValueType) (ValueType, error)
	Concat(vt ValueType) (ValueType, error)
	Sub(st ValueType) (ValueType, error)
	Mul(st ValueType) (ValueType, error)
	Div(st ValueType) (ValueType, error)
	Mod(st ValueType) (ValueType, error)
	BitAnd(st ValueType) (ValueType, error)
	BitOr(st ValueType) (ValueType, error)
	BitXOR(st ValueType) (ValueType, error)
	Equals(st ValueType) (ValueType, error)
	NotEquals(st ValueType) (ValueType, error)
	GreaterThan(st ValueType) (ValueType, error)
	LesserThan(st ValueType) (ValueType, error)
}
