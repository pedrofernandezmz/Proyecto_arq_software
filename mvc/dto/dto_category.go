package dto

type CategoryDto struct {
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
}

type CategoriesDto []CategoryDto
