package model

type ArticleCreateRequest struct {
	AuthorID string   `json:"author_id" validate:"required"`
	Title    string   `json:"title" validate:"required"`
	Status   string   `json:"status" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	TagNames []string `json:"tag_names" validate:"required,min=1,dive,required"`
}

type ArticleUpdateRequest struct {
	Slug     string   `json:"slug" validate:"required"`
	Title    string   `json:"title" validate:"required"`
	Status   string   `json:"status" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	TagNames []string `json:"tag_names" validate:"required,min=1,dive,required"`
}

type ArticleCreateResponse struct {
	Title    string   `json:"title" validate:"required"`
	Status   string   `json:"status" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	TagNames []string `json:"tag_names" validate:"required,min=1,dive,required"`
	Version  int64    `json:"version" validate:"required"`
}
