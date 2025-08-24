package service

import (
	"context"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/validation"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo example_domain.Repository
}

func New(repo example_domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, req model.ExampleCreateRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("example.service.Create: failed to validate request: %w", err)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		err = fmt.Errorf("device.service.Create: failed to generate id: %w", err)
		return
	}

	data := entity.Example{
		ID:          id.String(),
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	}

	err = s.repo.Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("example.service.Create: failed to create example: %w", err)
		return
	}

	return
}

func (s Service) GetByID(ctx context.Context, id string) (res model.ExampleResponse, err error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("example.service.GetByID: failed to get example: %w", err)
		return
	}
	res = model.ExampleResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Type:        data.Type,
		CreatedAt:   data.CreatedAt.Format(time.DateTime),
		UpdatedAt:   data.UpdatedAt.Format(time.DateTime),
	}

	return
}

func (s Service) Update(ctx context.Context, req model.ExampleUpdateRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("example.service.Update: failed to validate request: %w", err)
		return
	}

	data := entity.Example{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	}

	err = s.repo.Update(ctx, data)
	if err != nil {
		err = fmt.Errorf("example.service.Update: failed to update example: %w", err)
		return
	}

	return
}

func (s Service) Delete(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		err = fmt.Errorf("example.service.Delete: failed to delete example: %w", err)
		return
	}

	return
}
