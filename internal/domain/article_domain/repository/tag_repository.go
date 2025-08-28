package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) CreateTag(ctx context.Context, data entity.Tag) (*entity.Tag, error) {
	var tag entity.Tag
	err := r.db.QueryRow(ctx, `
        INSERT INTO tags (name, usage_count) 
        VALUES ($1, $2)
        RETURNING id, name, usage_count, created_at`,
		data.Name, data.UsageCount).Scan(
		&tag.ID, &tag.Name, &tag.UsageCount, &tag.CreatedAt)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("tag.repository.Create: failed to create tag: %w", err)
		return nil, err
	}

	return &tag, nil
}

func (r Repository) GetByNameTag(ctx context.Context, name string) (data entity.Tag, err error) {
	query := `
		SELECT id, name, usage_count, created_at
		FROM tags
		WHERE name = $1
	`

	err = r.db.QueryRow(ctx, query, name).Scan(&data.ID, &data.Name, &data.UsageCount, &data.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return data, constant.ErrTagNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				return data, constant.ErrTagNotFound
			}
		}
		err = fmt.Errorf("tag.repository.GetByName: failed to get tag: %w", err)
		return
	}

	return
}

func (r Repository) UpdateTag(ctx context.Context, data entity.Tag) error {
	cmd, err := r.db.Exec(ctx, `
		UPDATE tags
		SET usage_count = usage_count + 1, last_used_at = $1
		WHERE name = $2
	`, data.LastUsedAt, data.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrTagNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrTagNotFound
			}
		}
		err = fmt.Errorf("tag.repository.Update: failed to update tag: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrTagNotFound
		err = fmt.Errorf("tag.repository.Update: failed to update tag: %w", err)
		return err
	}

	return nil
}

func (r Repository) DecrementTag(ctx context.Context, data entity.Tag) error {
	cmd, err := r.db.Exec(ctx, `
       UPDATE tags
       SET usage_count = GREATEST(usage_count - 1, 0)
       WHERE name = $1
	`, data.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrTagNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrTagNotFound
			}
		}
		err = fmt.Errorf("tag.repository.Update: failed to update tag: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrTagNotFound
		err = fmt.Errorf("tag.repository.Update: failed to update tag: %w", err)
		return err
	}

	return nil
}

func (r Repository) GetAllTags(ctx context.Context) ([]entity.Tag, error) {
	query := `
		SELECT 
			id,
			name,
			usage_count,
			last_used_at,
			created_at
		FROM tags
		ORDER BY name DESC `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.GetAllTags: failed to query tags: %w", err)
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
			return nil, fmt.Errorf("repository.GetAllTags: failed to scan tag: %w", err)
		}

		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository.GetAllTags: iteration error: %w", err)
	}

	return tags, nil
}

func (r Repository) GetTagByID(ctx context.Context, id string) (data entity.Tag, err error) {
	query := `
		SELECT id, name, usage_count, created_at
		FROM tags
		WHERE id = $1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&data.ID, &data.Name, &data.UsageCount, &data.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return data, constant.ErrTagNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				return data, constant.ErrTagNotFound
			}
		}
		err = fmt.Errorf("tag.repository.GetByName: failed to get tag: %w", err)
		return
	}

	return
}
