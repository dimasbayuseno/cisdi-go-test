package user_domain

import (
	"context"
	user_repository "github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)
	WithTX(tx pgx.Tx) *user_repository.Repository
	Create(ctx context.Context, data entity.User) error
	GetByID(ctx context.Context, id string) (data entity.User, err error)
	Update(ctx context.Context, data entity.User) error
	Delete(ctx context.Context, id string) error
	GetByUsername(ctx context.Context, username string) (data entity.User, err error)
}
