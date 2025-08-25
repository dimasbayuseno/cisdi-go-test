package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	dbpostgres "github.com/dimasbayuseno/cisdi-go-test/pkg/db/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	db dbpostgres.Queryer
}

func New(db dbpostgres.Queryer) *Repository {
	return &Repository{db: db}
}

func (r Repository) Create(ctx context.Context, data entity.Article) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO articles (author_id, title, slug, status) 
		VALUES ($1, $2, $3, $4)`,
		data.AuthorID, data.Title, data.Slug, data.Status)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("article.repository.Create: failed to create article: %w", err)
		return err

	}
	return nil
}
