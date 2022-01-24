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
	fmt.Println(arr)
}
