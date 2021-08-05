package view_models

import "github.com/l5coder/ecommerce-microservice/product-service/domain/models"

type CategoryVm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewCategoryVm(model *models.Category) *CategoryVm {
	return &CategoryVm{
		ID:   model.Id(),
		Name: model.Name(),
		Slug: model.Slug(),
	}
}
