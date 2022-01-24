package cp

func Copy[T any](src *T) *T {
	cur := *src
	return &cur
}
