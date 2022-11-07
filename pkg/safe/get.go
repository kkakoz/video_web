package safe

func Get[T any](f func() T) T {
	defer func() {
		recover()
	}()
	return f()
}

func GetDef[T any](f func() T, defValue T) (t T) {
	defer func() {
		recover()
		t = defValue
	}()
	return f()
}

func MustGet[T *any, V any](t T, f func(T) V) V {
	if t != nil {
		return f(t)
	}
	return *new(V)
}
