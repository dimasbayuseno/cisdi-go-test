package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

// ========================================== ARTICLE ========================================== //

type Article struct {
	ID          uuid.UUID  `json:"id"`
	AuthorID    uuid.UUID  `json:"author_id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	PublishedAt *time.Time `json:"published_at"`
}

type ArticleStatus string

type GetBy string

const (
	ArticleStatusPublished ArticleStatus = "published"
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusArchived  ArticleStatus = "archived"

	GetBySlug GetBy = "slug"
	GetByID   GetBy = "id"
)

func (Article) TableName() string { return "articles" }

func IsArticleStatusValid(status string) bool {
	switch ArticleStatus(status) {
	case ArticleStatusPublished, ArticleStatusDraft, ArticleStatusArchived:
		return true
	default:
		return false
	}
}

// ========================================== ARTICLE VERSION ========================================== //

type ArticleVersion struct {
	ID                          uuid.UUID       `json:"id"`
	ArticleID                   uuid.UUID       `json:"article_id"`
	VersionNumber               int64           `json:"version_number"`
	Content                     string          `json:"content"`
	TrendingScore               decimal.Decimal `json:"trending_score"`
	CreatedAt                   time.Time       `json:"created_at"`
	ArticleTagRelationshipScore decimal.Decimal `json:"article_tag_relationship_score"`
}

func (ArticleVersion) TableName() string { return "article_versions" }

// ========================================== TAG ========================================== //

type Tag struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	UsageCount int64     `json:"usage_count"`
	LastUsedAt time.Time `json:"last_used_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Tag) TableName() string { return "tags" }

// ========================================== TAG CO-OCCURRENCE ========================================== //

type TagCooccurrence struct {
	TagAID            uuid.UUID `json:"tag_a_id"`
	TagBID            uuid.UUID `json:"tag_b_id"`
	CooccurrenceCount int       `json:"cooccurrence_count"`
}

func (TagCooccurrence) TableName() string { return "tag_cooccurrence" }

// ========================================== ARTICLE VERSION TAG ========================================== //

type ArticleVersionTag struct {
	ArticleVersionID uuid.UUID `json:"article_version_id"`
	TagID            uuid.UUID `json:"tag_id"`
}

func (ArticleVersionTag) TableName() string { return "article_version_tags" }
