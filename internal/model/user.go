package model

type UserCreateRequest struct {
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"password_hash" validate:"required"`
	Role         string `json:"role" validate:"required"`
	FullName     string `json:"full_name" validate:"required"`
}

type UserResponse struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	FullName     string `json:"full_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UserUpdateRequest struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	UpdatedAt    string `json:"updated_at"`
}

type GetByUsernameRequest struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
