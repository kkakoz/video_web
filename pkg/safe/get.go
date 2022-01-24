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
