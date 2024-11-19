//go:build integration

package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	repo "src/internal/repository"
	"src/internal/repository/utils"
)

type OrderRepoSuite struct {
	suite.Suite

	orderRepo repo.IOrderRepository
	orderID   int
	userID    int
}

func (o *OrderRepoSuite) TestOrderRepoUpdate(t provider.T) {
	t.Title("[Update] Update order")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update order", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		order, err := o.orderRepo.GetOrderByID(ctx, o.orderID)
		sCtx.Assert().NoError(err)

		order.Status = model.OrderStatusDone

		sCtx.WithNewParameters("ctx", ctx, "request", order)

		err = o.orderRepo.Update(ctx, order)
		sCtx.Assert().NoError(err)
	})
}

func (o *OrderRepoSuite) TestOrderRepoGetAllOrders(t provider.T) {
	t.Title("[GetAllOrders] Get all orders")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get all orders", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := utils.OrderBuilder{}.ToListAllOrders([]string{"total_price"})

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		orders, err := o.orderRepo.GetAllOrders(ctx, request)

		sCtx.Assert().NotEmpty(orders)
		sCtx.Assert().NoError(err)
	})
}

func (o *OrderRepoSuite) TestOrderRepoGetOrderByID(t provider.T) {
	t.Title("[GetOrderByID] Get order by id")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get order by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "request", o.orderID)

		order, err := o.orderRepo.GetOrderByID(ctx, o.orderID)

		sCtx.Assert().NotEmpty(order)
		sCtx.Assert().NoError(err)
	})
}

func (o *OrderRepoSuite) TestOrderRepoGetMyOrders(t provider.T) {
	t.Title("[GetMyOrders] Get my orders")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get my orders", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "request", o.userID)

		orders, err := o.orderRepo.GetMyOrders(ctx, o.userID)

		sCtx.Assert().NotEmpty(orders)
		sCtx.Assert().NoError(err)
	})
}
