package models

// ExampleModel usually used to return a GET  request or to be parsed during a PUT request
// For PATCH, usually updates 1 (seldom more) so just use params if needed
type ExampleModel struct {
	ID       uint   `json:"id" validate:"required,number,gte=1" example:"1"`
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Password string `json:"password" validate:"required,min=8" example:"password"`
	Optional string `json:"optional" validate:"omitempty" example:"optional field"` // Omitempty in the validation tag will only validate if it's provided
}

type ExampleModelCreate struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Password string `json:"password" validate:"required,min=8" example:"password"`
	Optional string `json:"optional" validate:"omitempty" example:"optional field"` // Omitempty in the validation tag will only validate if it's provided
}
