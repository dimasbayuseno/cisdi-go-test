package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/validation"
	"github.com/shopspring/decimal"
	"time"
)

type Service struct {
	repo article_domain.Repository
}

func New(repo article_domain.Repository) *Service {
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
		return
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
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()
	res, err := s.repo.WithTX(tx).Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("article.service.Create: failed to create article: %w", err)
		return
	}

	dataArticleVersion := entity.ArticleVersion{
		ArticleID:                   res.ID,
		VersionNumber:               1,
		Content:                     req.Content,
		TrendingScore:               decimal.Zero,
		ArticleTagRelationshipScore: decimal.Zero,
	}

	articleVersion, err := s.repo.WithTX(tx).CreateArticleVersion(ctx, dataArticleVersion)
	if err != nil {
		err = fmt.Errorf("article_version.service.Create: failed to create article_version: %w", err)
		return
	}
	for _, tagName := range req.TagNames {
		var dataTag entity.Tag
		existingTag, err := s.repo.WithTX(tx).GetByNameTag(ctx, tagName)
		if err != nil && errors.Is(err, constant.ErrTagNotFound) {
			newTag := entity.Tag{
				Name:       tagName,
				UsageCount: 1,
			}
			createdTag, err := s.repo.WithTX(tx).CreateTag(ctx, newTag)
			if err != nil {
				return fmt.Errorf("failed to create tag %s: %w", tagName, err)
			}
			dataTag = *createdTag
		} else if err != nil {
			return fmt.Errorf("failed to get tag %s: %w", tagName, err)
		} else {
			existingTag.LastUsedAt = time.Now()
			err = s.repo.UpdateTag(ctx, existingTag)
			if err != nil {
				return fmt.Errorf("failed to update tag %s: %w", tagName, err)
			}
			dataTag = existingTag
		}
		err = s.repo.WithTX(tx).CreateArticleVersionTag(ctx, articleVersion.ID, dataTag.ID)
		if err != nil {
			return fmt.Errorf("failed to create article version tag for %s: %w", tagName, err)
		}
	}
	tx.Commit(ctx)
	return
}

func (s Service) GetDetailArticleBySlug(ctx context.Context, slug string) (response model.ArticleDetailResponse, err error) {
	article, err := s.repo.GetArticleBySlug(ctx, slug)
	if err != nil {
		return
	}
	articleVersion, err := s.repo.GetLastArticleVersionNumber(ctx, article.ID)
	if err != nil {
		return
	}

	tags, err := s.repo.GetTagsByArticleVersionID(ctx, articleVersion.ID)
	var tagResponses []model.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, model.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	response = model.ArticleDetailResponse{
		ID:            article.ID,
		AuthorID:      article.AuthorID,
		Title:         article.Title,
		Slug:          article.Slug,
		Status:        article.Status,
		Content:       articleVersion.Content,
		PublishedAt:   article.PublishedAt,
		VersionNumber: articleVersion.VersionNumber,
		Tags:          tagResponses,
	}
	return
}
