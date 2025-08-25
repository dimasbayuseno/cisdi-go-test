package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/validation"
	"github.com/shopspring/decimal"
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

	if !entity.IsArticleStatusValid(req.Status) {
		err = constant.ErrInvalidStatusArticle
	}

	data := entity.Article{
		AuthorID:    req.AuthorID,
		Title:       req.Title,
		Slug:        pkgutil.CreateSlug(req.Title),
		Status:      string(entity.ArticleStatusDraft),
		PublishedAt: nil,
	}

	if req.Status == string(entity.ArticleStatusPublished) {
		data.PublishedAt = func() *time.Time { t := time.Now(); return &t }()
	}

	tx, err := s.repo.BeginTransaction(ctx)
	if err != nil {
		err = constant.ErrFailedTx
	}

	res, err := s.repo.WithTX(tx).Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("article.service.Create: failed to create article: %w", err)
		tx.Rollback(ctx)
		return
	}

	dataArticleVersion := entity.ArticleVersion{
		ArticleID:                   res.ID,
		VersionNumber:               1,
		Content:                     req.Content,
		TrendingScore:               decimal.Zero,
		ArticleTagRelationshipScore: decimal.Zero,
	}

	err = s.repo.WithTX(tx).CreateArticleVersion(ctx, dataArticleVersion)
	if err != nil {
		err = fmt.Errorf("article_version.service.Create: failed to create article_version: %w", err)
		tx.Rollback(ctx)
		return
	}
	for _, tagName := range req.TagNames {
		var dataTag entity.Tag
		existingTag, err := s.repo.GetByNameTag(ctx, tagName)
		if err != nil && errors.Is(err, constant.ErrTagNotFound) {
			newTag := entity.Tag{
				Name:       tagName,
				UsageCount: 1,
			}
			createdTag, err := s.repo.CreateTag(ctx, newTag)
			if err != nil {
				tx.Rollback(ctx)
				return fmt.Errorf("failed to create tag %s: %w", tagName, err)
			}
			dataTag = *createdTag
		} else if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to get tag %s: %w", tagName, err)
		} else {
			existingTag.LastUsedAt = time.Now()
			err = s.repo.UpdateTag(ctx, existingTag)
			if err != nil {
				tx.Rollback(ctx)
				return fmt.Errorf("failed to update tag %s: %w", tagName, err)
			}
			dataTag = existingTag
		}
		err = s.repo.CreateArticleVersionTag(ctx, dataTag)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to update tag %s: %w", tagName, err)
		}
	}

	return
}

//func (s Service) Create(ctx context.Context, req model.ArticleCreateRequest) (err error) {
