package dto

type CreateCategoryRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
