package model

import (
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type ArticleCreateRequest struct {
	AuthorID uuid.UUID `json:"author_id" validate:"required"`
	Title    string    `json:"title" validate:"required"`
	Status   string    `json:"status" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	TagNames []string  `json:"tag_names" validate:"required,min=1,dive,required"`
}

type ArticleUpdateRequest struct {
	ArticleID string   `json:"article_id" validate:"required"`
	AuthorID  string   `json:"author_id" validate:"required"`
	Slug      string   `json:"slug" validate:"required"`
	Title     string   `json:"title" validate:"required"`
	Status    string   `json:"status" validate:"required"`
	Content   string   `json:"content" validate:"required"`
	TagNames  []string `json:"tag_names" validate:"required,min=1,dive,required"`
}

type ArticleCreateResponse struct {
	Title    string   `json:"title" validate:"required"`
	Status   string   `json:"status" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	TagNames []string `json:"tag_names" validate:"required,min=1,dive,required"`
	Version  int64    `json:"version" validate:"required"`
}

type GetArticlesRequest struct {
	Status    string
	AuthorID  uuid.UUID
	TagID     uuid.UUID
	Page      int
	Limit     int
	SortBy    string
	SortOrder string
}

type TagResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type VersionContent struct {
	ID            uuid.UUID     `json:"id"`
	Content       string        `json:"content"`
	VersionNumber int64         `json:"version_number"`
	Tag           []TagResponse `json:"tags"`
}

type ArticleResponse struct {
	ID                          uuid.UUID       `json:"id"`
	AuthorID                    uuid.UUID       `json:"author_id"`
	Title                       string          `json:"title"`
	Slug                        string          `json:"slug"`
	Status                      string          `json:"status"`
	CreatedAt                   time.Time       `json:"created_at"`
	UpdatedAt                   time.Time       `json:"updated_at"`
	PublishedAt                 *time.Time      `json:"published_at"`
	VersionNumber               int64           `json:"version_number"`
	ArticleTagRelationshipScore decimal.Decimal `json:"article_tag_relationship_score"`
	Tags                        []TagResponse   `json:"tags"`
}

type ArticleDetailResponse struct {
	ID                          uuid.UUID       `json:"id"`
	AuthorID                    uuid.UUID       `json:"author_id"`
	Title                       string          `json:"title"`
	Slug                        string          `json:"slug"`
	Status                      string          `json:"status"`
	Content                     string          `json:"content"`
	CreatedAt                   time.Time       `json:"created_at"`
	UpdatedAt                   time.Time       `json:"updated_at"`
	PublishedAt                 *time.Time      `json:"published_at"`
	VersionNumber               int64           `json:"version_number"`
	ArticleTagRelationshipScore decimal.Decimal `json:"article_tag_relationship_score"`
	Tags                        []TagResponse   `json:"tags"`
}

type GetArticlesResponse struct {
	Articles   []ArticleResponse          `json:"articles"`
	Pagination pkgutil.PaginationResponse `json:"pagination"`
}

type AllArticleResponse struct {
	ID                          uuid.UUID        `json:"id"`
	AuthorID                    uuid.UUID        `json:"author_id"`
	Title                       string           `json:"title"`
	Slug                        string           `json:"slug"`
	Status                      string           `json:"status"`
	CreatedAt                   time.Time        `json:"created_at"`
	UpdatedAt                   time.Time        `json:"updated_at"`
	PublishedAt                 *time.Time       `json:"published_at"`
	ArticleTagRelationshipScore decimal.Decimal  `json:"article_tag_relationship_score"`
	Content                     []VersionContent `json:"tags"`
}
