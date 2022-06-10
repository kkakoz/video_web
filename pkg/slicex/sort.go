package slicex

func Sort[T any](slice []T, f func(v1 T, v2 T) bool) {
	for i := 0; i < len(slice); i++ {
		ele := slice[i]
		var j int
		for j = i; j > 0 && f(ele, slice[j-1]); j-- {
			slice[j] = slice[j-1]
		}
		slice[j] = ele
	}
}

// func Add[T any](slice []T, index int, value ...T) []T {
// 	next := slice[index:]
// 	newSlice := append(slice[:index], value...)
// 	return append(newSlice, next...)
// }
//
// func Remove[T any](slice []T, index int) []T {
// 	return append(slice[:index], slice[index+1:]...)
// }
