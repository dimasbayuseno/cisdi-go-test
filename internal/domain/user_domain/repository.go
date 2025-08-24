package user_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data entity.User) error
	GetByID(ctx context.Context, id string) (data entity.User, err error)
	Update(ctx context.Context, data entity.User) error
	Delete(ctx context.Context, id string) error
	GetByUsername(ctx context.Context, username string) (data entity.User, err error)
}
