package slicex

import "sort"

func Sort[T any](slice []T, f func(v1 T, v2 T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return f(slice[i], slice[j])
	})
}
