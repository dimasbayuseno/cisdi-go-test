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

func (r Repository) CreateArticleVersionTag(ctx context.Context, articleVersionID uuid.UUID, tagID uuid.UUID) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO article_version_tags (article_version_id, tag_id) 
		VALUES ($1, $2)`,
		articleVersionID, tagID)

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

func (r Repository) GetTagsByArticleVersionID(ctx context.Context, versionID uuid.UUID) ([]entity.Tag, error) {
	query := `
        SELECT t.id, t.name, t.usage_count, t.last_used_at, t.created_at
        FROM tags t
        JOIN article_version_tags avt ON t.id = avt.tag_id
        WHERE avt.article_version_id = $1
    `

	rows, err := r.db.Query(ctx, query, versionID)
	if err != nil {
		return nil, fmt.Errorf("repository.GetTagsByArticleVersionID: failed to query tags: %w", err)
	}
	defer rows.Close()

	var tags []entity.Tag
	for rows.Next() {
		var tag entity.Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.UsageCount,
			&tag.LastUsedAt,
			&tag.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("repository.GetTagsByArticleVersionID: failed to scan tag: %w", err)
		}
		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository.GetTagsByArticleVersionID: error iterating rows: %w", err)
	}

	return tags, nil
}
