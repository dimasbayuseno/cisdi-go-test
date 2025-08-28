package article_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, req model.ArticleCreateRequest) (err error)
	GetDetailArticleBySlug(ctx context.Context, slug string) (response model.ArticleDetailResponse, err error)
	GetArticles(ctx context.Context, userID uuid.UUID, userRole string, req model.GetArticlesRequest) (*model.GetArticlesResponse, error)
	UpdateArticle(ctx context.Context, id, status, userID, role string) (err error)
	DeleteArticle(ctx context.Context, id, userID, role string) (err error)
	CreateNewArticleVersion(ctx context.Context, req model.ArticleUpdateRequest) (err error)
	GetArticleVersions(ctx context.Context, articleID uuid.UUID, role, userID string) (versions model.AllArticleResponse, err error)
	GetDetailArticleVersion(ctx context.Context, id string, versionNumber int) (response model.ArticleDetailResponse, err error)
	CreateNewTagByAdmin(ctx context.Context, name, role string) error
	GetAllTags(ctx context.Context, role string) ([]entity.Tag, error)
	GetDetailTag(ctx context.Context, id string, role string) (response entity.Tag, err error)
}
