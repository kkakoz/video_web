package lox

func Group[T comparable, V any](values []V, f func(value V) T) map[T]V {
	m := map[T]V{}
	for _, v := range values {
		m[f(v)] = v
	}
	return m
}
