//go:build unit

package service_test

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
)

type CartServiceSuite struct {
	suite.Suite
}

// AddRacket
func (s *CartServiceSuite) TestCartServiceAddRacket1(t provider.T) {
	t.Title("[AddRacket] No racket")
	t.Tags("cart", "service", "add_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: no racket", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1}.AddCartRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).AddRacket(ctx, req)

		sCtx.Assert().Nil(cart)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *CartServiceSuite) TestCartServiceAddRacket2(t provider.T) {
	t.Title("[AddRacket] Create cart")
	t.Tags("cart", "service", "add_racket")
	t.Parallel()
	t.WithNewStep("Success: create cart", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.AddCartRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		racket := &model.Racket{
			ID:        req.RacketID,
			Avaliable: true,
			Quantity:  100,
			Price:     100,
		}

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(racket, nil).
			Once()

		expCart := &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price),
			Quantity:   req.Quantity,
			Lines: []*model.CartLine{
				{
					RacketID: racket.ID,
					Quantity: req.Quantity,
					Price:    float32(racket.Price),
				},
			},
		}

		cartMockRepo.
			On("Create", ctx, expCart).
			Return(nil).
			Once()

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).AddRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}

func (s *CartServiceSuite) TestCartServiceAddRacket3(t provider.T) {
	t.Title("[AddRacket] add new racket")
	t.Tags("cart", "service", "add_racket")
	t.Parallel()
	t.WithNewStep("Success: add new racket", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.AddCartRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		expCart := &model.Cart{
			UserID:     req.UserID,
			TotalPrice: 0,
			Quantity:   0,
			Lines:      nil,
		}

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(expCart, nil).
			Once()

		racket := &model.Racket{
			ID:        req.RacketID,
			Avaliable: true,
			Quantity:  100,
			Price:     100,
		}

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(racket, nil).
			Once()

		cartMockRepo.
			On("AddRacket", ctx, req).
			Return(nil).
			Once()

		cartMockRepo.
			On("Update", ctx, expCart).
			Return(nil).
			Once()

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).AddRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
		sCtx.Assert().Equal(cart, expCart)
	})
}

func (s *CartServiceSuite) TestCartServiceRemoveRacket1(t provider.T) {
	t.Title("[RemoveRacket] no cart yet")
	t.Tags("cart", "service", "remove_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: no cart yet", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.RemoveRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		expCart := &model.Cart{
			UserID: req.UserID,
		}

		cartMockRepo.
			On("Create", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).RemoveRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}

func (s *CartServiceSuite) TestCartServiceRemoveRacket2(t provider.T) {
	t.Title("[RemoveRacket] remove not existed racket")
	t.Tags("cart", "service", "remove_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: remove not existed racket", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.RemoveRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		racket := &model.Racket{
			ID:    req.RacketID,
			Price: 100,
		}

		expCart := &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price),
			Quantity:   1,
			Lines: []*model.CartLine{
				{
					RacketID: racket.ID,
					Quantity: 1,
				},
			},
		}

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(expCart, nil).
			Once()

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).RemoveRacket(ctx, req)

		sCtx.Assert().Nil(cart)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *CartServiceSuite) TestCartServiceRemoveRacket3(t provider.T) {
	t.Title("[RemoveRacket] remove existed racket")
	t.Tags("cart", "service", "remove_racket")
	t.Parallel()
	t.WithNewStep("Success: remove existed racket", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.RemoveRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		racket := &model.Racket{
			ID:       req.RacketID,
			Price:    100,
			Quantity: 100,
		}

		expCart := &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price),
			Quantity:   1,
			Lines: []*model.CartLine{
				{
					RacketID: req.RacketID,
					Quantity: 1,
					Price:    float32(racket.Price),
				},
			},
		}

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(expCart, nil).
			Once()

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(racket, nil).
			Once()

		cartMockRepo.
			On("RemoveRacket", ctx, req).
			Return(nil).
			Once()

		cartMockRepo.
			On("Update", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).RemoveRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}

func (s *CartServiceSuite) TestCartServiceUpdateRacket1(t provider.T) {
	t.Title("[UpdateRacket] no cart yet")
	t.Tags("cart", "service", "update_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: no cart yet", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.UpdatePlusRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(nil, fmt.Errorf("get cart by id fail")).
			Once()

		expCart := &model.Cart{
			UserID: req.UserID,
		}

		cartMockRepo.
			On("Create", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).UpdateRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}

func (s *CartServiceSuite) TestCartServiceUpdateRacket2(t provider.T) {
	t.Title("[UpdateRacket] add racket quantity")
	t.Tags("cart", "service", "update_racket")
	t.Parallel()
	t.WithNewStep("Success: add racket quantity", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.UpdatePlusRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		racket := &model.Racket{
			ID:       req.RacketID,
			Price:    100,
			Quantity: 100,
		}

		expCart := &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price),
			Quantity:   1,
			Lines: []*model.CartLine{
				{
					RacketID: req.RacketID,
					Quantity: 1,
					Price:    float32(racket.Price),
				},
			},
		}

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(expCart, nil).
			Once()

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(racket, nil).
			Once()

		cartMockRepo.
			On("Update", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).UpdateRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
		sCtx.Assert().Equal(cart.Quantity, 2)
	})
}

func (s *CartServiceSuite) TestCartServiceUpdateRacket3(t provider.T) {
	t.Title("[UpdateRacket] subtract racket quantity")
	t.Tags("cart", "service", "update_racket")
	t.Parallel()
	t.WithNewStep("Success: subtract racket quantity", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1, Quantity: 1}.UpdateRacketMinusReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		racket := &model.Racket{
			ID:       req.RacketID,
			Price:    100,
			Quantity: 100,
		}

		expCart := &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price),
			Quantity:   1,
			Lines: []*model.CartLine{
				{
					RacketID: req.RacketID,
					Quantity: 1,
					Price:    float32(racket.Price),
				},
			},
		}

		cartMockRepo.
			On("GetCartByID", ctx, req.UserID).
			Return(expCart, nil).
			Once()

		racketMockRepo.
			On("GetRacketByID", ctx, req.RacketID).
			Return(racket, nil).
			Once()

		cartMockRepo.
			On("Update", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).UpdateRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
		sCtx.Assert().Equal(cart.Quantity, 0)
	})
}

// GetCartByID
func (s *CartServiceSuite) TestCartServiceGetCartByID1(t provider.T) {
	t.Title("[GetCartByID] no existed cart")
	t.Tags("cart", "service", "get_cart_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: no existed cart", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1}.GetCartByID()

		cartMockRepo := mocks.NewICartRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx, req).
			Return(nil, fmt.Errorf("get cart by id fail")).
			Once()

		expCart := &model.Cart{
			UserID: req,
		}

		cartMockRepo.
			On("Create", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, nil).GetCartByID(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}

func (s *CartServiceSuite) TestCartServiceGetCartByID2(t provider.T) {
	t.Title("[GetCartByID] existed cart")
	t.Tags("cart", "service", "get_cart_by_id")
	t.Parallel()
	t.WithNewStep("Success: existed cart", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1}.GetCartByID()

		cartMockRepo := mocks.NewICartRepository(t)

		expCart := &model.Cart{
			UserID:     req,
			TotalPrice: 0,
			Quantity:   0,
			Lines:      nil,
		}

		cartMockRepo.
			On("GetCartByID", ctx, req).
			Return(expCart, nil).
			Once()

		cartMockRepo.
			On("Update", ctx, expCart).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, nil).GetCartByID(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}
