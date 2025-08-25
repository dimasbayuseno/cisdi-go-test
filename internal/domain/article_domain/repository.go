package article_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
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
}
