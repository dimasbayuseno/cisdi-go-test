package model

type ExampleCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

type ExampleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ExampleUpdateRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
