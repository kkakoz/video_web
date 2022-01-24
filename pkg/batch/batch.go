package batch

func Filter[T any](data []T, f func(T) bool) []T {
	res := []T{}
	for _, v := range data {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
