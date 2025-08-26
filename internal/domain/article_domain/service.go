package article_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
)

type Service interface {
	Create(ctx context.Context, req model.ArticleCreateRequest) (err error)
	GetDetailArticleBySlug(ctx context.Context, slug string) (response model.ArticleDetailResponse, err error)
}
