package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	dbpostgres "github.com/dimasbayuseno/cisdi-go-test/pkg/db/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	db dbpostgres.Queryer
}

func New(db dbpostgres.Queryer) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Create(ctx context.Context, data entity.Example) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO examples (id, name, description, type)
		VALUES ($1, $2, $3, $4)
	`, data.ID, data.Name, data.Description, data.Type)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrDeviceNotFound
			}
		}

		err = fmt.Errorf("example.repository.Create: failed to create example: %w", err)
		return err
	}

	return nil
}

func (r Repository) GetByID(ctx context.Context, id string) (data entity.Example, err error) {
	query := `
		SELECT id, name, description, type, created_at, updated_at
		FROM examples
		WHERE id = $1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&data.ID, &data.Name, &data.Description, &data.Type, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrExampleNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrExampleNotFound
			}
		}
		err = fmt.Errorf("example.repository.GetByID: failed to get example: %w", err)
		return
	}

	return
}

func (r Repository) Update(ctx context.Context, data entity.Example) error {
	cmd, err := r.db.Exec(ctx, `
		UPDATE examples
		SET name = $1, description = $2, type = $3
		WHERE id = $4
	`, data.Name, data.Description, data.Type, data.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrExampleNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrExampleNotFound
			}
		}
		err = fmt.Errorf("example.repository.Update: failed to update example: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrExampleNotFound
		err = fmt.Errorf("device.repository.Update: failed to update device: %w", err)
		return err
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id string) error {
	cmd, err := r.db.Exec(ctx, `
		DELETE FROM examples
		WHERE id = $1
	`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrExampleNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrExampleNotFound
			}
		}
		err = fmt.Errorf("example.repository.Delete: failed to delete example: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrExampleNotFound
		err = fmt.Errorf("device.repository.Update: failed to update device: %w", err)
		return err
	}

	return nil
}
