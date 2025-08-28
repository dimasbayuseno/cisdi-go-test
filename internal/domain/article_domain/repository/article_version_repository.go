package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) CreateArticleVersion(ctx context.Context, data entity.ArticleVersion) (res *entity.ArticleVersion, err error) {
	var result entity.ArticleVersion

	err = r.db.QueryRow(ctx, `
       INSERT INTO article_versions (article_id, version_number, content, trending_score) 
       VALUES ($1, $2, $3, $4)
       RETURNING id, article_id, version_number, content, trending_score, created_at`,
		data.ArticleID, data.VersionNumber, data.Content, data.TrendingScore).Scan(
		&result.ID,
		&result.ArticleID,
		&result.VersionNumber,
		&result.Content,
		&result.TrendingScore,
		&result.CreatedAt,
	)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("article_version.repository.Create: failed to create article version: %w", err)
		return nil, err
	}

	return &result, nil
}

func (r Repository) CreateNewArticleVersion(ctx context.Context, data entity.ArticleVersion, newVersionNumber int64) (res entity.ArticleVersion, err error) {
	var newVersionID uuid.UUID
	err = r.db.QueryRow(ctx, `
		INSERT INTO article_versions (article_id, version_number, content) 
		VALUES ($1, $2, $3) RETURNING id`,
		data.ArticleID, newVersionNumber, data.Content).Scan(&newVersionID)

	if err != nil {
		return res, fmt.Errorf("article_version.repository.Create: failed to create article version: %w", err)
	}

	data.ID = newVersionID
	data.VersionNumber = newVersionNumber

	return data, nil
}

func (r Repository) GetLastArticleVersionNumber(ctx context.Context, articleID uuid.UUID) (data entity.ArticleVersion, err error) {
	err = r.db.QueryRow(ctx, `SELECT id, article_id, version_number, content FROM article_versions WHERE article_id = $1 ORDER BY version_number DESC LIMIT 1`, articleID).Scan(&data.ID, &data.ArticleID, &data.VersionNumber, &data.Content)
	if err != nil {
		return data, fmt.Errorf("repository.GetLastArticleVersionNumber: failed to get last version number: %w", err)
	}

	return data, nil
}
