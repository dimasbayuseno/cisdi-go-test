package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
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
