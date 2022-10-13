package slicex

import "golang.org/x/exp/constraints"

func Add[T any](arr []T, index int, value ...T) []T {
	cur := Clone(arr)
	next := append(value, cur[index:]...)
	cur = append(cur[:index], next...)
	return cur
}

func Delete[T any](arr []T, index int) []T {
	cur := Clone(arr)
	before := cur[:index]
	cur = append(before, cur[index+1:]...)
	return cur
}

func Clone[T any](arr []T) []T {
	newData := make([]T, len(arr))
	copy(newData, arr)
	return newData
}

func SliceMap[T any, K constraints.Ordered, V any](s []T, f func(v T) (K, V)) map[K]V {

	m := map[K]V{}
	for _, v := range s {
		k, v := f(v)
		m[k] = v
	}
	return m
}

func SafeGet[T any](arr []T, index int) T {
	if index > len(arr) {
		return *new(T)
	}
	return arr[index]
}
