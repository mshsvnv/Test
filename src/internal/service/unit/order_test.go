//go:build unit
package service_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
)

type OrderServiceSuite struct {
	suite.Suite
}

// CreateOrder
func (s *OrderServiceSuite) TestOrderServiceCreateOrder1(t provider.T) {
	t.Title("[CreateOrder1] no existed cart")
	t.Tags("order", "service", "create_order")
	t.Parallel()
	t.WithNewStep("Incorrect: no existed cart", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{
			UserID:        1,
			DeliveryDate:  time.Now().AddDate(0, 1, 1),
			Address:       "Moscow",
			RecepientName: "Stepan Postnov",
			Status:        model.OrderStatusInProgress,
			Lines: []*model.OrderLine{
				{
					RacketID: 1,
					Quantity: 1,
				},
			},
			TotalPrice: 100,
		}.ToPlaceOrderDTO()

		cartMockRepo := mocks.NewICartRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(nil, fmt.Errorf("get cart fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := service.NewOrderService(utils.NewMockLogger(), nil, cartMockRepo, nil).CreateOrder(ctx, req)

		sCtx.Assert().Error(err)
	})
}

func (s *OrderServiceSuite) TestOrderServiceCreateOrder2(t provider.T) {
	t.Title("[CreateOrder2] create order")
	t.Tags("order", "service", "create_order")
	t.Parallel()
	t.WithNewStep("Success: create order", func(sCtx provider.StepCtx) {

		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		ctx := context.TODO()
		orderBuilder := utils.OrderBuilder{
			UserID:        1,
			DeliveryDate:  time.Now().AddDate(0, 1, 1),
			Address:       "Moscow",
			RecepientName: "Stepan Postnov",
			CreationDate:  tm,
			Status:        model.OrderStatusInProgress,
		}

		racket := utils.RacketObjectMother{}.DefaultRacket()
		cart := utils.CartObjectMother{
			UserID:   orderBuilder.OrderID,
			RacketID: racket.ID,
			Quantity: 1,
		}.DefaultCart()

		order := orderBuilder.ToModel(cart)
		req := orderBuilder.ToPlaceOrderDTO()

		orderMockRepo := mocks.NewIOrderRepository(t)
		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(cart, nil).
			Once()

		racketMockRepo.
			On("GetRacketByID", ctx, racket.ID).
			Return(racket, nil).
			Once()

		racketMockRepo.
			On("Update", ctx, racket).
			Return(nil).
			Once()

		orderMockRepo.
			On("Create", ctx, order).
			Return(nil).
			Once()

		cartMockRepo.
			On("Delete", ctx, req.UserID).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, cartMockRepo, racketMockRepo).CreateOrder(ctx, req)

		sCtx.Assert().Nil(err)
		sCtx.Assert().Equal(racket.Quantity, 99)
	})
}

// GetMyOrders
func (s *OrderServiceSuite) TestOrderServiceGetMyOrders1(t provider.T) {
	t.Title("[GetMyOrders1] wrong user id")
	t.Tags("order", "service", "get_my_orders")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong user id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{UserID: 0}

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetMyOrders", ctx, req.UserID).
			Return(nil, fmt.Errorf("get my orders fail, error")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req.UserID)

		orders, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).GetMyOrders(ctx, req.UserID)

		sCtx.Assert().Empty(orders)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "get my orders fail, error")
	})
}

func (s *OrderServiceSuite) TestOrderServiceGetMyOrders2(t provider.T) {
	t.Title("[GetMyOrders2] correct user id")
	t.Tags("order", "service", "get_my_orders")
	t.Parallel()
	t.WithNewStep("Success: correct user id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{}.WithUserID()

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetMyOrders", ctx, req.UserID).
			Return([]*model.Order{
				{
					UserID: req.UserID,
					Lines: []*model.OrderLine{
						{
							RacketID: 1,
							Quantity: 1,
						},
					},
				},
			}, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req.UserID)

		orders, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).GetMyOrders(ctx, req.UserID)

		sCtx.Assert().NotEmpty(orders)
		sCtx.Assert().NoError(err)
	})
}

// GetAllOrders
func (s *OrderServiceSuite) TestOrderServiceGetAllOrders1(t provider.T) {
	t.Title("[GetAllOrders1] get all orders")
	t.Tags("order", "service", "get_all_orders")
	t.Parallel()
	t.WithNewStep("Success: get all orders", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{}.ToListAllOrders([]string{"total_price"})

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetAllOrders", ctx, req).
			Return([]*model.Order{
				{
					TotalPrice: 100,
				},
				{
					TotalPrice: 200,
				},
			}, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		orders, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).GetAllOrders(ctx, req)

		sCtx.Assert().NotEmpty(orders)
		sCtx.Assert().NoError(err)
	})
}

func (s *OrderServiceSuite) TestOrderServiceGetAllOrders2(t provider.T) {
	t.Title("[GetAllOrders2] wrong column name")
	t.Tags("order", "service", "get_all_orders")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong column name", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{}.ToListAllOrders([]string{"totalprice"})

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetAllOrders", ctx, req).
			Return(nil, fmt.Errorf("get all fail, error")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		orders, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).GetAllOrders(ctx, req)

		sCtx.Assert().Nil(orders)
		sCtx.Assert().Error(err)
	})
}

// // GetOrderByID
func (s *OrderServiceSuite) TestOrderServiceGetOrderByID1(t provider.T) {
	t.Title("[GetOrderByID1] wrong order id")
	t.Tags("order", "service", "get_order_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong order id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{OrderID: 0}

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetOrderByID", ctx, req.OrderID).
			Return(nil, fmt.Errorf("get order by id")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req.OrderID)

		order, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).GetOrderByID(ctx, req.OrderID)

		sCtx.Assert().Empty(order)
		sCtx.Assert().Error(err)
	})
}

func (s *OrderServiceSuite) TestOrderServiceGetOrderByID2(t provider.T) {
	t.Title("[GetOrderByID2] correct order id")
	t.Tags("order", "service", "get_order_by_id")
	t.Parallel()
	t.WithNewStep("Success: correct order id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{}.WithOrderID()

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetOrderByID", ctx, req.OrderID).
			Return(&model.Order{
				UserID:     req.UserID,
				TotalPrice: 1,
			}, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req.OrderID)

		order, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).GetOrderByID(ctx, req.OrderID)

		sCtx.Assert().NotEmpty(order)
		sCtx.Assert().Nil(err)
	})
}

// UpdateOrderStatus
func (s *OrderServiceSuite) TestOrderServiceUpdateOrderStatus(t provider.T) {
	t.Title("[UpdateOrderStatus] correct order id")
	t.Tags("order", "service", "get_order_by_id")
	t.Parallel()
	t.WithNewStep("Success: correct order id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.OrderBuilder{}.WithOrderID().ToUpdateDTO(model.OrderStatusDone)
		orderTmp := &model.Order{
			ID:         req.OrderID,
			TotalPrice: 100,
		}

		orderMockRepo := mocks.NewIOrderRepository(t)
		orderMockRepo.
			On("GetOrderByID", ctx, req.OrderID).
			Return(orderTmp, nil).
			Once()

		orderTmp.Status = model.OrderStatusDone
		orderMockRepo.
			On("Update", ctx, orderTmp).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		order, err := service.NewOrderService(utils.NewMockLogger(), orderMockRepo, nil, nil).UpdateOrderStatus(ctx, req)

		sCtx.Assert().NotEmpty(order)
		sCtx.Assert().Nil(err)
	})
}
