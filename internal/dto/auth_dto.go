package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenClaimsRequest struct {
	UserID        string `json:"user_id"`
	DepartmentID  string `json:"department_id"`
	Role          string `json:"role"`
	Secret        string `json:"secret"`
	ExpiresMinute int    `json:"expires_minute"`
}

type RefreshTokenClaimsRequest struct {
	UserID       string `json:"user_id"`
	DepartmentID string `json:"department_id"`
	Role         string `json:"role"`
	Secret       string `json:"secret"`
	ExpiresDay   int    `json:"expires_day"`
}
