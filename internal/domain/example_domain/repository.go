package example_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data entity.Example) error
	GetByID(ctx context.Context, id string) (data entity.Example, err error)
	Update(ctx context.Context, data entity.Example) error
	Delete(ctx context.Context, id string) error
}
