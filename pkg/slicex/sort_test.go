package slicex

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{1, 5, 4, 2, 3}
	Sort(arr, func(v1 int, v2 int) bool {
		return v1 < v2
	})

}

func TestFilter(t *testing.T) {
	arr := []int{1, 5, 4, 2, 3}
	// time.Local = time.FixedZone("UTC+8", 8*3600)
	// lo.Filter(arr, func(v int, index int) bool {
	// 	return v%2 == 0
	// })
	// fmt.Println(arr)
	// lo.ForEach(arr, func(v int, index int) {
	// 	fmt.Println(lo.Ternary(v%2 == 0, v, 0))
	// })
	//
	//
	// // fmt.Println(Add(arr, 2, 2, 3, 4))
	// fmt.Println(Delete(arr, 1))
	// fmt.Println(Delete(arr, 1))
	// // fmt.Println(Remove(arr, 4))
	//
	// ti := time.Unix(1654494671, 0)

	// fmt.Println(ti)

	// arr = slices.Delete(arr, 1, 2)
	// fmt.Println(arr)
	//
	// fmt.Println(slices.Insert(arr, 4, 1))
	//
	// fmt.Println(arr)
	fmt.Println(Add(arr, 1, 10))
	fmt.Println(arr)

	fmt.Println(Delete(arr, 3))
	fmt.Println(arr)
}
