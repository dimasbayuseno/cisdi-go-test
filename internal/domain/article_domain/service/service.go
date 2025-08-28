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
	"github.com/google/uuid"
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
			err = s.repo.WithTX(tx).UpdateTag(ctx, existingTag)
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

func (s Service) GetArticles(ctx context.Context, userID uuid.UUID, userRole string, req model.GetArticlesRequest) (*model.GetArticlesResponse, error) {

	articles, err := s.repo.GetArticles(ctx, userRole, userID, req)
	if err != nil {
		return nil, fmt.Errorf("article.service.GetArticles: repository error: %w", err)
	}

	totalCount, err := s.repo.GetArticlesCount(ctx, userRole, userID, req)
	if err != nil {

		totalCount = 0
	}

	pagination := pkgutil.BuildPagination(req.Page, req.Limit, int64(totalCount))

	return &model.GetArticlesResponse{
		Articles:   articles,
		Pagination: pagination,
	}, nil
}

func (s Service) UpdateArticle(ctx context.Context, id, status, userID, role string) (err error) {
	article, err := s.repo.GetArticleByID(ctx, id)
	if err != nil {
		return
	}

	userIDNew, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %v", err)
	}

	switch role {
	case string(entity.RoleAdmin):
	case string(entity.RoleEditor):
		if userIDNew != article.AuthorID {
			err = constant.ErrUserNotMatch
			return
		}
	default:
		err = constant.ErrUnauthorizedAccess
		return
	}

	return s.repo.UpdateArticleStatusWithPublishDate(ctx, article.ID, status)
}

func (s Service) DeleteArticle(ctx context.Context, id, userID, role string) (err error) {
	article, err := s.repo.GetArticleByID(ctx, id)
	if err != nil {
		return
	}

	userIDNew, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %v", err)
	}

	switch role {
	case string(entity.RoleAdmin):
	case string(entity.RoleEditor):
		if userIDNew != article.AuthorID {
			err = constant.ErrUserNotMatch
			return
		}
	default:
		err = constant.ErrUnauthorizedAccess
		return
	}

	articleVersion, err := s.repo.GetLastArticleVersionNumber(ctx, article.ID)
	if err != nil {
		return
	}
	articleTags, err := s.repo.GetTagsByArticleVersionID(ctx, articleVersion.ID)
	if err != nil {
		return
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
	for _, tag := range articleTags {
		err = s.repo.WithTX(tx).DecrementTag(ctx, tag)
		if err != nil {
			return
		}
	}

	err = s.repo.WithTX(tx).Delete(ctx, article.ID)
	if err != nil {
		return
	}
	tx.Commit(ctx)
	return nil
}

func (s Service) CreateNewArticleVersion(ctx context.Context, req model.ArticleUpdateRequest) (err error) {
	article, err := s.repo.GetArticleByID(ctx, req.ArticleID)
	if err != nil {
		return
	}

	reqAuthorID, err := uuid.Parse(req.AuthorID)
	if err != nil {
		return
	}

	if reqAuthorID != article.AuthorID {
		err = constant.ErrUserNotMatch
		return
	}

	articleVersion, err := s.repo.GetLastArticleVersionNumber(ctx, article.ID)
	if err != nil {
		return
	}

	articleTags, err := s.repo.GetTagsByArticleVersionID(ctx, articleVersion.ID)
	if err != nil {
		return
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

	for _, tag := range articleTags {
		err = s.repo.WithTX(tx).DecrementTag(ctx, tag)
		if err != nil {
			return
		}
	}

	dataArticleVersion := entity.ArticleVersion{
		ArticleID:     article.ID,
		VersionNumber: articleVersion.VersionNumber + 1,
		Content:       req.Content,
		CreatedAt:     time.Now(),
	}

	newArticleVersion, err := s.repo.WithTX(tx).CreateNewArticleVersion(ctx, dataArticleVersion, articleVersion.VersionNumber+1)
	if err != nil {
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
			err = s.repo.WithTX(tx).UpdateTag(ctx, existingTag)
			if err != nil {
				return fmt.Errorf("failed to update tag %s: %w", tagName, err)
			}
			dataTag = existingTag
		}
		err = s.repo.WithTX(tx).CreateArticleVersionTag(ctx, newArticleVersion.ID, dataTag.ID)
		if err != nil {
			return fmt.Errorf("failed to create article version tag for %s: %w", tagName, err)
		}
	}

	return
}

func (s Service) GetArticleVersions(ctx context.Context, articleID uuid.UUID, role, userID string) (versions model.AllArticleResponse, err error) {

	article, err := s.repo.GetArticleByID(ctx, articleID.String())
	if err != nil {
		return
	}

	userIDNew, err := uuid.Parse(userID)
	if err != nil {
		return versions, fmt.Errorf("invalid user ID format: %v", err)
	}

	switch role {
	case string(entity.RoleAdmin):
	case string(entity.RoleEditor):
		if userIDNew != article.AuthorID {
			err = constant.ErrUserNotMatch
			return
		}
	default:
		err = constant.ErrUnauthorizedAccess
		return
	}

	articleVersions, err := s.repo.GetArticleVersions(ctx, articleID)
	if err != nil {
		return
	}

	var contents []model.VersionContent

	for _, version := range articleVersions {
		tags, err := s.repo.GetTagsByArticleVersionID(ctx, version.ID)
		if err != nil {
			return
		}

		var tagResponses []model.TagResponse
		for _, tag := range tags {
			tagResponses = append(tagResponses, model.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			})
		}

		contents = append(contents, model.VersionContent{
			ID:            version.ID,
			Content:       version.Content,
			VersionNumber: version.VersionNumber,
			Tag:           tagResponses,
		})
	}

	versions = model.AllArticleResponse{
		ID:          article.ID,
		AuthorID:    userIDNew,
		Title:       article.Title,
		Slug:        article.Slug,
		Status:      article.Status,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		PublishedAt: article.PublishedAt,
		Content:     contents,
	}

	return versions, nil

}

func (s Service) GetDetailArticleVersion(ctx context.Context, id string, versionNumber int) (response model.ArticleDetailResponse, err error) {
	article, err := s.repo.GetArticleByID(ctx, id)
	if err != nil {
		return
	}
	articleVersion, err := s.repo.GetArticleVersionByNumber(ctx, article.ID, versionNumber)
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

func (s Service) CreateNewTagByAdmin(ctx context.Context, name, role string) error {

	if role != string(entity.RoleAdmin) {
		return constant.ErrUnauthorizedAccess
	}

	newTag := entity.Tag{
		Name:       name,
		UsageCount: 0,
	}
	_, err := s.repo.CreateTag(ctx, newTag)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetAllTags(ctx context.Context, role string) ([]entity.Tag, error) {
	if role != string(entity.RoleAdmin) {
		return nil, constant.ErrUnauthorizedAccess
	}
	tags, err := s.repo.GetAllTags(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil

}

func (s Service) GetDetailTag(ctx context.Context, id string, role string) (response entity.Tag, err error) {
	if role != string(entity.RoleAdmin) {
		return response, constant.ErrUnauthorizedAccess
	}
	response, err = s.repo.GetTagByID(ctx, id)
	if err != nil {
		return
	}

	return
}
