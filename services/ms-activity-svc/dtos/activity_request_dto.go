package dtos

type ActivityRequestDTO struct {
	Name string `json:"name" binding:"required,min=4,max=33"`
}
