package dto

type CategoryAdd struct {
	Name string `json:"name" validate:"required"`
}
