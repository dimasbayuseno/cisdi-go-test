package service

import (
	"context"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/validation"
	"time"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, req model.ArticleCreateRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("article.service.Register: failed to validate request: %w", err)
		return
	}

	data := entity.Article{
		ID:          "",
		AuthorID:    "",
		Title:       "",
		Slug:        "",
		Status:      string(entity.ArticleStatusDraft),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		PublishedAt: nil,
	}

	err = s.repo.Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("article.service.Create: failed to create article: %w", err)
		return
	}

	return
}
