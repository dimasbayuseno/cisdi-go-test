package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	dbpostgres "github.com/dimasbayuseno/cisdi-go-test/pkg/db/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	db dbpostgres.Queryer
}

func New(db dbpostgres.Queryer) *Repository {
	return &Repository{db: db}
}

func (r Repository) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r Repository) WithTX(tx pgx.Tx) *Repository {
	return &Repository{db: tx}
}

func (r Repository) Create(ctx context.Context, data entity.Article) (*entity.Article, error) {
	var id uuid.UUID
	err := r.db.QueryRow(ctx, `
        INSERT INTO articles (author_id, title, slug, status) 
        VALUES ($1, $2, $3, $4)
        RETURNING id`,
		data.AuthorID, data.Title, data.Slug, data.Status).Scan(&id)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("article.repository.Create: failed to create article: %w", err)
		return nil, err
	}

	data.ID = id
	return &data, nil
}
