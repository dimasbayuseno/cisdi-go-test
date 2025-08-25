package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) CreateArticleVersionTag(ctx context.Context, data entity.Tag) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO article_version_tags (article_version_id, tag_id) 
		VALUES ($1, $2)`,
		data.Name, data.UsageCount)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("article_version_tag.repository.Create: failed to create article_version_tag: %w", err)
		return err

	}
	return nil
}
