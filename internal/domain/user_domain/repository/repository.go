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

func (r Repository) Create(ctx context.Context, data entity.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users (username, email, password_hash, role, full_name)
		VALUES ($1, $2, $3, $4, $5)
	`, data.Username, data.Email, data.PasswordHash, data.Role, data.FullName)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}

		err = fmt.Errorf("user.repository.Create: failed to create user: %w", err)
		return err
	}

	return nil
}

func (r Repository) GetByID(ctx context.Context, id string) (data entity.User, err error) {
	query := `
		SELECT id, username, email, password_hash, role, full_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&data.ID, &data.Username, &data.Email, &data.PasswordHash, &data.Role, &data.FullName, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrUserNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrUserNotFound
			}
		}
		err = fmt.Errorf("user.repository.GetByID: failed to get user: %w", err)
		return
	}

	return
}

func (r Repository) Update(ctx context.Context, data entity.User) error {
	cmd, err := r.db.Exec(ctx, `
		UPDATE users
		SET full_name = $1, role = $2, password_hash = $3, updated_at = $4
		WHERE id = $5
	`, data.FullName, data.Role, data.PasswordHash, data.UpdatedAt, data.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrUserNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrUserNotFound
			}
		}
		err = fmt.Errorf("user.repository.Update: failed to update user: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrUserNotFound
		err = fmt.Errorf("user.repository.Update: failed to update user: %w", err)
		return err
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id string) error {
	cmd, err := r.db.Exec(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrUserNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrUserNotFound
			}
		}
		err = fmt.Errorf("user.repository.Delete: failed to delete user: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrUserNotFound
		err = fmt.Errorf("user.repository.Update: failed to update user: %w", err)
		return err
	}

	return nil
}

func (r Repository) GetByUsername(ctx context.Context, username string) (data entity.User, err error) {
	query := `
		SELECT id, username, email, password_hash, role, full_name, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	err = r.db.QueryRow(ctx, query, username).Scan(&data.ID, &data.Username, &data.Email, &data.PasswordHash, &data.Role, &data.FullName, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrUserNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrUserNotFound
			}
		}
		err = fmt.Errorf("user.repository.GetByID: failed to get user: %w", err)
		return
	}

	return
}
