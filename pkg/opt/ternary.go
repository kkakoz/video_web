package opt

func Ternary[T any](b bool, v1, v2 T) T {
	if b {
		return v1
	}
	return v2
}

func IfThen[T any](b bool, v1, v2 T) T {
	if b {
		return v1
	}
	return v2
}
