package request

import "vidlearn-final-projcect/business/category/spec"

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func (request *CreateCategoryRequest) ToSpec() *spec.UpsertCategorySpec {
	return &spec.UpsertCategorySpec{
		Name:     request.Name,
		ImageURL: request.ImageURL,
	}
}
