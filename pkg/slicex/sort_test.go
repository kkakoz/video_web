package slicex

import (
	"fmt"
	"github.com/samber/lo"
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
	lo.Filter(arr, func(v int, index int) bool {
		return v%2 == 0
	})
	fmt.Println(arr)
	lo.ForEach(arr, func(v int, index int) {
		fmt.Println(lo.Ternary(v%2 == 0, v, 0))
	})
}
