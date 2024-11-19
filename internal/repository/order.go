package repository

import (
	"context"
	"src/internal/dto"
	"src/internal/model"
)

//go:generate mockery --name=IOrderRepository
type IOrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, orderID int) error
	GetAllOrders(ctx context.Context, req *dto.ListOrdersReq) ([]*model.Order, error)
	GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
}
