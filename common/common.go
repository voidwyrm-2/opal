package common

func Assert[T any](v T, _ error) T {
	return v
}

func AssertOk[T any](v T, _ bool) T {
	return v
}
