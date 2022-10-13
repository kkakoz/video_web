package mathx

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer | constraints.Float](v T) T {
	return -v
}

func Sum[T constraints.Integer | constraints.Float](nums []T) T {
	sum := *new(T)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max[T constraints.Integer | constraints.Float](v1 T, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min[T constraints.Integer | constraints.Float](v1 T, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}
