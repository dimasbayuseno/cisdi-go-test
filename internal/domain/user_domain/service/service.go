package service

import (
	"context"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/validation"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	repo user_domain.Repository
}

func New(repo user_domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, req model.UserCreateRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("user.service.Register: failed to validate request: %w", err)
		return
	}
	if !entity.IsRoleValid(req.Role) {
		err = fmt.Errorf("user.service.Register: wrong role: %w", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("user.service.Register: failed to hash password: %w", err)
		return
	}

	data := entity.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
		FullName:     req.FullName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.repo.Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("user.service.Create: failed to create user: %w", err)
		return
	}

	return
}

func (s Service) GetByID(ctx context.Context, id string) (res model.UserResponse, err error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("user.service.GetByID: failed to get user: %w", err)
		return
	}
	res = model.UserResponse{
		ID:           data.ID,
		Username:     data.Username,
		Email:        data.Email,
		PasswordHash: data.PasswordHash,
		Role:         data.Role,
		FullName:     data.FullName,
		CreatedAt:    data.CreatedAt.Format(time.DateTime),
		UpdatedAt:    data.UpdatedAt.Format(time.DateTime),
	}

	return
}

func (s Service) Update(ctx context.Context, req model.UserUpdateRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("user.service.Update: failed to validate request: %w", err)
		return
	}

	data := entity.User{
		ID:           req.ID,
		Username:     req.UpdatedAt,
		Email:        req.ID,
		PasswordHash: req.PasswordHash,
		Role:         req.Role,
		FullName:     req.FullName,
		UpdatedAt:    time.Now(),
	}

	err = s.repo.Update(ctx, data)
	if err != nil {
		err = fmt.Errorf("user.service.Update: failed to update user: %w", err)
		return
	}

	return
}

func (s Service) Delete(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		err = fmt.Errorf("user.service.Delete: failed to delete user: %w", err)
		return
	}

	return
}

func (s Service) Login(ctx context.Context, req model.LoginRequest) (res model.LoginResponse, err error) {
	data, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		err = fmt.Errorf("user.service.GetByUsername: failed to get user: %w", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(data.PasswordHash))
	if err != nil {
		err = fmt.Errorf("user.service.Login: wrong password: %w", err)
	}

	token, err := pkgutil.GenerateJWT(data.ID, data.Username, data.Email, data.Role)
	if err != nil {
		return res, fmt.Errorf("user.service.Login: failed to generate token: %w", err)
	}

	res = model.LoginResponse{
		Token: token,
	}

	return
}
