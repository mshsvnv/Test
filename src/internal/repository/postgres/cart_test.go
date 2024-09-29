package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/repository/mocks"
	"src/internal/repository/utils"
)

type CartRepoSuite struct {
	suite.Suite

	cartMockRepo mocks.ICartRepository
}

func (c *CartRepoSuite) BeforeAll(t provider.T) {
	t.Title("Init cart mock repo")
	c.cartMockRepo = *mocks.NewICartRepository(t)
	t.Tags("fixture", "cart")
}

func (c *CartRepoSuite) TestCartRepoCreate(t provider.T) {
	t.Title("[Create] Create cart")
	t.Tags("order repository", "postgres")
	t.Parallel()
	t.WithNewStep("Create cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CartObjectMother{
			UserID:   1,
			RacketID: 1,
			Quantity: 100,
		}.DefaultCart()

		c.cartMockRepo.
			On("Create", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartMockRepo.Create(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *CartRepoSuite) TestCartRepoAddRacket(t provider.T) {
	t.Title("[AddRacket] Add racket to cart")
	t.Tags("order repository", "postgres")
	t.Parallel()
	t.WithNewStep("Add racket to cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CartObjectMother{
			UserID:   1,
			RacketID: 1,
			Quantity: 100,
		}.AddCartRacketReq()

		c.cartMockRepo.
			On("AddRacket", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartMockRepo.AddRacket(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *CartRepoSuite) TestCartRepoRemoveRacket(t provider.T) {
	t.Title("[RemoveRacket] Remove racket to cart")
	t.Tags("order repository", "postgres")
	t.Parallel()
	t.WithNewStep("Remove racket to cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CartObjectMother{
			UserID:   1,
			RacketID: 1,
			Quantity: 100,
		}.RemoveRacketReq()

		c.cartMockRepo.
			On("RemoveRacket", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartMockRepo.RemoveRacket(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *CartRepoSuite) TestCartRepoUpdate(t provider.T) {
	t.Title("[Update] Update cart")
	t.Tags("order repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CartObjectMother{
			UserID:   1,
			RacketID: 1,
			Quantity: 100,
		}.DefaultCart()

		c.cartMockRepo.
			On("Update", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartMockRepo.Update(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *CartRepoSuite) TestCartRepoGetCartByID(t provider.T) {
	t.Title("[GetCartByID] Get cart by id")
	t.Tags("order repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get cart by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := 1
		expCart := utils.CartObjectMother{
			UserID:   request,
			RacketID: 1,
			Quantity: 100,
		}.DefaultCart()

		c.cartMockRepo.
			On("GetCartByID", ctx, request).
			Return(expCart, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		cart, err := c.cartMockRepo.GetCartByID(ctx, request)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(cart, expCart)
	})
}

func (c *CartRepoSuite) Delete(t provider.T) {
	t.Title("[Delete] Get cart by id")
	t.Tags("order repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get cart by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := 1
		c.cartMockRepo.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartMockRepo.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
