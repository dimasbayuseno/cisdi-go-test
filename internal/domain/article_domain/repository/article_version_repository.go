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

func (r Repository) CreateArticleVersion(ctx context.Context, data entity.ArticleVersion) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO article_versions (article_id, version_number, content, trending_score) 
		VALUES ($1, $2, $3, $4)`,
		data.ArticleID, data.VersionNumber, data.Content, data.TrendingScore)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("article_version.repository.Create: failed to create article version: %w", err)
		return err

	}
	return nil
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

	data.VersionNumber = newVersionNumber

	return res, nil
}

func (r Repository) GetLastArticleVersionNumber(ctx context.Context, articleID uuid.UUID) (data entity.ArticleVersion, err error) {
	err = r.db.QueryRow(ctx, `SELECT id, article_id, version_number, content FROM article_versions WHERE article_id = $1`, articleID).Scan(&data.ID, &data.ArticleID, &data.VersionNumber, &data.Content)
	if err != nil {
		return data, fmt.Errorf("repository.GetLastArticleVersionNumber: failed to get last version number: %w", err)
	}

	return data, nil
}
