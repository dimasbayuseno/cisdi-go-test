package user_domain

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
)

type Service interface {
	Create(ctx context.Context, req model.UserCreateRequest) (err error)
	GetByID(ctx context.Context, id string) (res model.UserResponse, err error)
	Update(ctx context.Context, req model.UserUpdateRequest) (err error)
	Delete(ctx context.Context, id string) (err error)
	Login(ctx context.Context, req model.LoginRequest) (res model.LoginResponse, err error)
}
