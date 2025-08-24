package example_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
)

type Service interface {
	Create(ctx context.Context, req model.ExampleCreateRequest) (err error)
	GetByID(ctx context.Context, id string) (res model.ExampleResponse, err error)
	Update(ctx context.Context, req model.ExampleUpdateRequest) (err error)
	Delete(ctx context.Context, id string) (err error)
}
