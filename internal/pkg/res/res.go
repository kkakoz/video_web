package res

type More[T any] struct {
	More  bool
	Value []*T
}

func MoreList[T any](list []*T, limit int) More[T] {
	res := More[T]{
		Value: list[:limit],
	}
	if len(list) > limit {
		res.More = true
	}
	return res
}
