package response

import "vidlearn-final-projcect/business/category"

type GetCategoryByIDResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func CreateGetAllCategoryResponse(categories []*category.Category) []*GetCategoryByIDResponse {
	var categoriesResponse []*GetCategoryByIDResponse

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, CreateGetCategoryByIDResponse(category))
	}

	return categoriesResponse
}

func CreateGetCategoryByIDResponse(category *category.Category) *GetCategoryByIDResponse {
	return &GetCategoryByIDResponse{
		ID:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}
}
