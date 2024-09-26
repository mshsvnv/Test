package service_test

import (
	"context"
	"fmt"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
)

type CartSuite struct {
	suite.Suite
}

// AddRacket
func (s *CartSuite) TestAddRacket(t provider.T) {
	t.Title("[AddRacket] Success")
	t.Tags("cart", "add_racket")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.CartObjectMother{UserID: 1, RacketID: 1}.AddCartRacketReq()

		cartMockRepo := mocks.NewICartRepository(t)
		racketMockRepo := mocks.NewIRacketRepository(t)

		cartMockRepo.
			On("GetCartByID", ctx).
			Return(nil, fmt.Errorf("get cart fails, cart doesn't exist")).
			Once()

		// racket := &model.Racket{
		// 	ID: req.RacketID,
		// }

		// racketMockRepo.
		// 	On("GetRacketByID", ctx, 1).
		// 	Return(racket, nil).
		// 	Once()

		// cartMockRepo.
		// 	On("Update", ctx, expCart).
		// 	Return(nil).
		// 	Once()

		// sCtx.WithNewParameters("ctx", ctx, "request", req)

		cart, err := service.NewCartService(utils.NewMockLogger(), cartMockRepo, racketMockRepo).AddRacket(ctx, req)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Nil(err)
	})
}

// func (s *AuthSuite) TestRemoveRacket(t provider.T) {
// 	t.Title("[RemoveRacket] User already exists")
// 	t.Tags("cart", "remove_racket")
// 	t.Parallel()
// 	t.WithNewStep("Incorrect: user already exists", func(sCtx provider.StepCtx) {

// 		ctx := context.TODO()
// 		req := utils.AuthObjectMother{}.DefaultUserReq()

// 		sCtx.WithNewParameters("ctx", ctx, "request", req)

// 		token, err := s.authService.Register(ctx, req)

// 		sCtx.Assert().Empty(token)
// 		sCtx.Assert().Error(err)
// 	})
// }

// func (s *AuthSuite) TestUpdateRacket(t provider.T) {
// 	t.Title("[RemoveRacket] User already exists")
// 	t.Tags("cart", "remove_racket")
// 	t.Parallel()
// 	t.WithNewStep("Incorrect: user already exists", func(sCtx provider.StepCtx) {

// 		ctx := context.TODO()
// 		req := utils.AuthObjectMother{}.DefaultUserReq()

// 		sCtx.WithNewParameters("ctx", ctx, "request", req)

// 		token, err := s.authService.Register(ctx, req)

// 		sCtx.Assert().Empty(token)
// 		sCtx.Assert().Error(err)
// 	})
// }

// func (s *AuthSuite) TestUpdateRacket(t provider.T) {
// 	t.Title("[UpdateRacket] User already exists")
// 	t.Tags("cart", "update_racket")
// 	t.Parallel()
// 	t.WithNewStep("Incorrect: user already exists", func(sCtx provider.StepCtx) {

// 		ctx := context.TODO()
// 		req := utils.AuthObjectMother{}.DefaultUserReq()

// 		sCtx.WithNewParameters("ctx", ctx, "request", req)

// 		token, err := s.authService.Register(ctx, req)

// 		sCtx.Assert().Empty(token)
// 		sCtx.Assert().Error(err)
// 	})
// }

// func (s *AuthSuite) TestGetCartByID(t provider.T) {
// 	t.Title("[GetCartByID] User already exists")
// 	t.Tags("cart", "get_cart_by_id")
// 	t.Parallel()
// 	t.WithNewStep("Incorrect: user already exists", func(sCtx provider.StepCtx) {

// 		ctx := context.TODO()
// 		req := utils.AuthObjectMother{}.DefaultUserReq()

// 		sCtx.WithNewParameters("ctx", ctx, "request", req)

// 		token, err := s.authService.Register(ctx, req)

// 		sCtx.Assert().Empty(token)
// 		sCtx.Assert().Error(err)
// 	})
// }
