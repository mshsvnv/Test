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
	})
}

func (c *CartRepoSuite) TestCartRepoRemoveRacket(t provider.T) {
	t.Title("[RemoveRacket] Remove racket to cart")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Remove racket to cart", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		reqAdd := utils.CartObjectMother{
			UserID:   c.cartID,
			RacketID: c.racketID,
			Quantity: 1,
		}.AddCartRacketReq()
		err := c.cartRepo.AddRacket(ctx, reqAdd)

		reqRemove := utils.CartObjectMother{
			UserID:   c.cartID,
			RacketID: c.racketID,
			Quantity: 1,
		}.RemoveRacketReq()

		sCtx.WithNewParameters("ctx", ctx, "request", reqRemove)

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
		sCtx.Assert().Equal(cart.Quantity, 0)
		sCtx.Assert().Equal(cart.TotalPrice, float32(0))
	})
}

func (c *CartRepoSuite) TestCartRepoDelete(t provider.T) {
	t.Title("[Delete] Delete cart by id")
	t.Tags("integration", "order", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Delete cart by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		user := utils.UserObjectMother{}.
			WithName("Masha").
			WithEmail("mshsvnv").
			ToModel()
		err := c.userRepo.Create(ctx, user)
		sCtx.Assert().NoError(err)

		cart := utils.CartObjectMother{
			UserID:   user.ID,
			Quantity: 1,
			RacketID: c.racketID,
		}.DefaultCart()

		err = c.cartRepo.Create(ctx, cart)
		sCtx.Assert().NoError(err)

		sCtx.WithNewParameters("ctx", ctx, "request", user.ID)

		err = c.cartRepo.Delete(ctx, user.ID)
		sCtx.Assert().NoError(err)
	})
}
