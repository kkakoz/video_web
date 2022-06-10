package jsonx

import "encoding/json"

func Unmarshal[T any](data []byte) (*T, error) {
	value := new(T)
	err := json.Unmarshal(data, value)
	return value, err
}
