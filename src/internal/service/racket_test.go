package service_test

import (
	"context"
	"fmt"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
)

type RacketSuite struct {
	suite.Suite

	racketService service.IRacketService
}

// CreateRacket
func (s *RacketSuite) TestCreateRacket1(t provider.T) {
	t.Title("[CreateRacket] wrong amount")
	t.Tags("racket", "create_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong amount", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.IncorrectCount()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := s.racketService.CreateRacket(ctx, req)

		sCtx.Assert().Nil(racket)
		sCtx.Assert().Error(err)
	})
}

func (s *RacketSuite) TestCreateRacket2(t provider.T) {
	t.Title("[CreateRacket] correct request")
	t.Tags("racket", "create_racket")
	t.Parallel()
	t.WithNewStep("Success: correct request", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.CorrectCount()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := s.racketService.CreateRacket(ctx, req)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().Nil(err)
	})
}

// UpdateRacket
func (s *RacketSuite) TestUpdateRacket1(t provider.T) {
	t.Title("[UpdateRacket] wrong ID")
	t.Tags("racket", "update_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong ID", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.UpdateIncorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := s.racketService.UpdateRacket(ctx, req)

		sCtx.Assert().Error(err)
	})
}

func (s *RacketSuite) TestUpdateRacket2(t provider.T) {
	t.Title("[UpdateRacket] correct request")
	t.Tags("racket", "update_racket")
	t.Parallel()
	t.WithNewStep("Success: correct request", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.UpdateCorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := s.racketService.UpdateRacket(ctx, req)

		sCtx.Assert().Nil(err)
	})
}

// GetRacketByID
func (s *RacketSuite) TestGetRacketByID1(t provider.T) {
	t.Title("[GetRacketByID] wrong id")
	t.Tags("racket", "get_racket_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.GetIncorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := s.racketService.GetRacketByID(ctx, req)

		sCtx.Assert().Nil(racket)
		sCtx.Assert().Error(err)
	})
}

func (s *RacketSuite) TestGetRacketByID2(t provider.T) {
	t.Title("[GetRacketByID] correct id")
	t.Tags("racket", "get_racket_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: correct id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.GetCorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := s.racketService.GetRacketByID(ctx, req)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().Nil(err)
	})
}

// GetAllRackets
func (s *RacketSuite) TestGetAllRackets1(t provider.T) {
	t.Title("[GetAllRackets] no rackets")
	t.Tags("racket", "get_all_rackets")
	t.Parallel()
	t.WithNewStep("Incorrect: no rackets", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.IncorrectFieldToSort()

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetAllRackets", ctx, req).
			Return(nil, fmt.Errorf("get all fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		rackets, err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).GetAllRackets(ctx, req)

		sCtx.Assert().Nil(rackets)
		sCtx.Assert().Error(err)
	})
}

func (s *RacketSuite) TestGetAllRackets2(t provider.T) {
	t.Title("[GetAllRackets] get rackets")
	t.Tags("racket", "get_all_rackets")
	t.Parallel()
	t.WithNewStep("Success: get rackets", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.CorrectFieldToSort()

		rackets := []*model.Racket{
			{
				ID:    1,
				Brand: "head",
			},
			{
				ID:    2,
				Brand: "babolat",
			},
		}

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetAllRackets", ctx, req).
			Return(rackets, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		rackets, err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).GetAllRackets(ctx, req)

		sCtx.Assert().NotEmpty(rackets)
		sCtx.Assert().Nil(err)
	})
}
