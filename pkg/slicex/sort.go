package slicex

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func Sort[T any](slice []T, f func(v1 T, v2 T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return f(slice[i], slice[j])
	})
}

func SortNum[T constraints.Integer | constraints.Float](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
