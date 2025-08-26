package article_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)
	WithTX(tx pgx.Tx) *repository.Repository
	Create(ctx context.Context, data entity.Article) (*entity.Article, error)
	CreateArticleVersion(ctx context.Context, data entity.ArticleVersion) error
	CreateTag(ctx context.Context, data entity.Tag) (*entity.Tag, error)
	GetByNameTag(ctx context.Context, name string) (data entity.Tag, err error)
	UpdateTag(ctx context.Context, data entity.Tag) error
	DecrementTag(ctx context.Context, data entity.Tag) error
	CreateNewArticleVersion(ctx context.Context, data entity.ArticleVersion, newVersionNumber int64) (res entity.ArticleVersion, err error)
	GetArticles(ctx context.Context, role string, currentUserID uuid.UUID, params model.GetArticlesRequest) ([]model.ArticleResponse, error)
	GetArticleDetails(ctx context.Context, articleID uuid.UUID, role string, currentUserID uuid.UUID) (*model.ArticleDetailResponse, error)
	GetLastArticleVersionNumber(ctx context.Context, articleID uuid.UUID) (data entity.ArticleVersion, err error)
	GetArticleBySlug(ctx context.Context, slug string) (data entity.Article, err error)
	GetTagsByArticleVersionID(ctx context.Context, versionID uuid.UUID) ([]entity.Tag, error)
}
