package dto

type CreateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	DepartmentID string `json:"department_id"`
	Role         string `json:"role"`
}

type UpdateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	DepartmentID string `json:"department_id"`
	Role         string `json:"role"`
}

type UserResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	DepartmentID string `json:"department_id"`
	Role         string `json:"role"`
}
