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

func (r Repository) CreateTag(ctx context.Context, data entity.Tag) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO tags (name, usage_count) 
		VALUES ($1, $2)`,
		data.Name, data.UsageCount)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("tag.repository.Create: failed to create tag: %w", err)
		return err

	}
	return nil
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
			err = constant.ErrTagNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrTagNotFound
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
		SET usage_count = usage_count + 1
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
