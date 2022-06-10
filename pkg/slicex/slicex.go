package slicex

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
