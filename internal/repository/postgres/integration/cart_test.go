//go:build integration

package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	repo "src/internal/repository"
	"src/internal/repository/utils"
)

type CartRepoSuite struct {
	suite.Suite

	cartRepo repo.ICartRepository
	userRepo repo.IUserRepository
	cartID   int
	racketID int
}

func (c *CartRepoSuite) TestCartRepoCreate(t provider.T) {
	t.Title("[Create] Create cart")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Create cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request, _ := c.cartRepo.GetCartByID(ctx, c.cartID)

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartRepo.Create(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (c *CartRepoSuite) TestCartRepoAddRacket(t provider.T) {
	t.Title("[AddRacket] Add racket to cart")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Add racket to cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CartObjectMother{
			UserID:   c.cartID,
			RacketID: c.racketID,
			Quantity: 1,
		}.AddCartRacketReq()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.cartRepo.AddRacket(ctx, request)
		sCtx.Assert().NoError(err)

		reqRemove := utils.CartObjectMother{
			UserID:   c.cartID,
			RacketID: c.racketID - 1,
			Quantity: 1,
		}.RemoveRacketReq()

		err = c.cartRepo.RemoveRacket(ctx, reqRemove)
		sCtx.Assert().NoError(err)
	})
}

func (c *CartRepoSuite) TestCartRepoGetCartByID(t provider.T) {
	t.Title("[GetCartByID] Get cart by id")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get cart by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		request := c.cartID

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		cart, err := c.cartRepo.GetCartByID(ctx, request)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().NoError(err)
	})
}
