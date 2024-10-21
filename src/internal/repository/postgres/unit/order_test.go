//go:build unit

package mypostgres_test

import (
	"context"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/repository/utils"
)

type OrderRepoSuite struct {
	suite.Suite

	orderMockRepo mocks.IOrderRepository
}

func (o *OrderRepoSuite) BeforeAll(t provider.T) {
	t.Title("Init order mock repo")
	o.orderMockRepo = *mocks.NewIOrderRepository(t)
	t.Tags("fixture", "order")
}

func (o *OrderRepoSuite) TestOrderRepoCreate(t provider.T) {
	t.Title("[Create] Create order")
	t.Tags("order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Create order", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		cart := utils.CartObjectMother{
			RacketID: 1,
			Quantity: 1,
			UserID:   1,
		}.DefaultCart()
		request := utils.OrderBuilder{
			UserID:        1,
			CreationDate:  tm,
			DeliveryDate:  tm,
			RecepientName: "Stepan Postnov",
			Address:       "Moscow",
			Status:        model.OrderStatusInProgress,
			TotalPrice:    0,
		}.ToModel(cart)

		o.orderMockRepo.
			On("Create", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := o.orderMockRepo.Create(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (o *OrderRepoSuite) TestOrderRepoUpdate(t provider.T) {
	t.Title("[Update] Update order")
	t.Tags("order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update order", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		cart := utils.CartObjectMother{
			RacketID: 1,
			Quantity: 1,
			UserID:   1,
		}.DefaultCart()
		request := utils.OrderBuilder{
			UserID:        1,
			CreationDate:  tm,
			DeliveryDate:  tm,
			RecepientName: "Stepan Postnov",
			Address:       "Moscow",
			Status:        model.OrderStatusDone,
			TotalPrice:    0,
		}.ToModel(cart)

		o.orderMockRepo.
			On("Update", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := o.orderMockRepo.Update(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (o *OrderRepoSuite) TestOrderRepoDelete(t provider.T) {
	t.Title("[Delete] Update order")
	t.Tags("order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Delete order", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := 1

		o.orderMockRepo.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := o.orderMockRepo.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (o *OrderRepoSuite) TestOrderRepoGetAllOrders(t provider.T) {
	t.Title("[GetAllOrders] Get all orders")
	t.Tags("order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get all orders", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		ordersID := []int{1}
		expOrders := []*model.Order{
			{
				ID:            1,
				UserID:        1,
				CreationDate:  tm,
				DeliveryDate:  tm,
				RecepientName: "Stepan Postnov",
				Address:       "Moscow",
			},
		}
		request := utils.OrderBuilder{}.ToListAllOrders([]string{"total_price"})

		o.orderMockRepo.
			On("GetAllOrders", ctx, request).
			Return(expOrders, nil).
			Once()

		o.orderMockRepo.
			On("getAllOrders", ctx, request).
			Return(ordersID, nil).
			Once()

		o.orderMockRepo.
			On("GetOrderByID", ctx, ordersID[0]).
			Return(expOrders[0], nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		orders, err := o.orderMockRepo.GetAllOrders(ctx, request)

		sCtx.Assert().NotEmpty(orders)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expOrders, orders)
	})
}

func (o *OrderRepoSuite) TestOrderRepoGetOrderByID(t provider.T) {
	t.Title("[GetOrderByID] Get order by id")
	t.Tags("order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get order by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := utils.OrderBuilder{OrderID: 1}.OrderID
		tm, _ := time.Parse(time.RFC3339, "2006-01-02")
		expOrder := &model.Order{
			ID:            request,
			UserID:        1,
			CreationDate:  tm,
			DeliveryDate:  tm,
			RecepientName: "Stepan Postnov",
			Address:       "Moscow",
		}

		o.orderMockRepo.
			On("GetOrderByID", ctx, request).
			Return(expOrder, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		order, err := o.orderMockRepo.GetOrderByID(ctx, request)

		sCtx.Assert().NotEmpty(order)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expOrder, order)
	})
}

func (o *OrderRepoSuite) TestOrderRepoGetMyOrders(t provider.T) {
	t.Title("[GetMyOrders] Get my orders")
	t.Tags("order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get my orders", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := utils.OrderBuilder{UserID: 1}.UserID
		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		ordersID := []int{1}
		expOrders := []*model.Order{
			{
				ID:            ordersID[0],
				UserID:        request,
				CreationDate:  tm,
				DeliveryDate:  tm,
				RecepientName: "Stepan Postnov",
				Address:       "Moscow",
			},
		}

		o.orderMockRepo.
			On("GetMyOrders", ctx, request).
			Return(expOrders, nil).
			Once()

		o.orderMockRepo.
			On("getMyOrders", ctx, request).
			Return(ordersID, nil).
			Once()

		o.orderMockRepo.
			On("GetOrderByID", ctx, ordersID[0]).
			Return(expOrders[0], nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		orders, err := o.orderMockRepo.GetMyOrders(ctx, request)

		sCtx.Assert().NotEmpty(orders)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expOrders, orders)
	})
}
