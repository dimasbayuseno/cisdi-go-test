package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Article struct {
	ID          string     `json:"id"`
	AuthorID    string     `json:"author_id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	PublishedAt *time.Time `json:"published_at"`
}

type ArtictleStatus string

const (
	ArtictleStatusPublished ArtictleStatus = "published"
	ArticleStatusDraft      ArtictleStatus = "draft"
	ArticleStatusArchived   ArtictleStatus = "archived"
)

func (Article) TableName() string { return "articles" }

func IsArticleStatusValid(status string) bool {
	switch ArtictleStatus(status) {
	case ArtictleStatusPublished, ArticleStatusDraft, ArticleStatusArchived:
		return true
	default:
		return false
	}
}

// ========================================== ARTICLE VERSION ========================================== //

type ArticleVersion struct {
	ID            string          `json:"id"`
	ArticleID     string          `json:"article_id"`
	VersionNumber string          `json:"version_number"`
	Content       string          `json:"content"`
	TrendingScore decimal.Decimal `json:"trending_score"`
	CreatedAt     time.Time       `json:"created_at"`
}

func (ArticleVersion) TableName() string { return "article_versions" }

// ========================================== TAG ========================================== //

type Tag struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	UsageCount string    `json:"usage_count"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Tag) TableName() string { return "tags" }
