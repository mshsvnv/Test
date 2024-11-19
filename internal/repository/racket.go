package repository

import (
	"context"
	"src/internal/dto"
	"src/internal/model"
)

//go:generate mockery --name=IRacketRepository
type IRacketRepository interface {
	Create(ctx context.Context, racket *model.Racket) error
	Update(ctx context.Context, racket *model.Racket) error
	Delete(ctx context.Context, id int) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error)
}
