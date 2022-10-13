package stringx

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

func ToInteger[T constraints.Integer](str string) (T, error) {
	v, err := strconv.Atoi(str)
	return T(v), err
}

func MustInteger[T constraints.Integer](str string) T {
	v, _ := strconv.Atoi(str)
	return T(v)
}
