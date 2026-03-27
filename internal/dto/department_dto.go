package dto

type CreateDepartmentRequest struct {
	Name string `json:"name"`
}

type UpdateDepartmentRequest struct {
	Name string `json:"name"`
}

type DepartmentResponse struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	UserResponses *[]UserResponse `json:"users,omitempty"`
}
