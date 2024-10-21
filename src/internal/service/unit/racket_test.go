//go:build unit

package service_test

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/dto"
	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
	utils1 "src/pkg/utils"
)

type RacketServiceSuite struct {
	suite.Suite
}

// CreateRacket
func (s *RacketServiceSuite) TestRacketServiceCreateRacket1(t provider.T) {
	t.Title("[CreateRacket] wrong amount")
	t.Tags("racket", "service", "create_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong amount", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.IncorrectCount()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := service.NewRacketService(utils.NewMockLogger(), nil).CreateRacket(ctx, req)

		sCtx.Assert().Nil(racket)
		sCtx.Assert().Error(err)
	})
}

func (s *RacketServiceSuite) TestRacketServiceCreateRacket2(t provider.T) {
	t.Title("[CreateRacket] correct request")
	t.Tags("racket", "service", "create_racket")
	t.Parallel()
	t.WithNewStep("Success: correct request", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		racket := utils.RacketObjectMother{}.DefaultRacket()
		var req dto.CreateRacketReq

		utils1.Copy(&req, racket)

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("Create", ctx, racket).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).CreateRacket(ctx, &req)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().NoError(err)
	})
}

// UpdateRacket
func (s *RacketServiceSuite) TestRacketServiceUpdateRacket1(t provider.T) {
	t.Title("[UpdateRacket] wrong ID")
	t.Tags("racket", "service", "update_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong ID", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.UpdateIncorrectID()

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetRacketByID", ctx, req.ID).
			Return(nil, fmt.Errorf("get racket by id")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).UpdateRacket(ctx, req)

		sCtx.Assert().Error(err)
	})
}

func (s *RacketServiceSuite) TestRacketServiceUpdateRacket2(t provider.T) {
	t.Title("[UpdateRacket] correct request")
	t.Tags("racket", "service", "update_racket")
	t.Parallel()
	t.WithNewStep("Success: correct request", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.UpdateCorrectID()
		racket := utils.RacketObjectMother{}.DefaultRacket()

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetRacketByID", ctx, req.ID).
			Return(racket, nil).
			Once()

		racketMockRepo.
			On("Update", ctx, racket).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).UpdateRacket(ctx, req)

		sCtx.Assert().NoError(err)
	})
}

// GetRacketByID
func (s *RacketServiceSuite) TestRacketServiceGetRacketByID1(t provider.T) {
	t.Title("[GetRacketByID] wrong id")
	t.Tags("racket", "service", "get_racket_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.GetIncorrectID()

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetRacketByID", ctx, req).
			Return(nil, fmt.Errorf("get racket by id fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).GetRacketByID(ctx, req)

		sCtx.Assert().Nil(racket)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "get racket by id fail")
	})
}

func (s *RacketServiceSuite) TestRacketServiceGetRacketByID2(t provider.T) {
	t.Title("[GetRacketByID] correct id")
	t.Tags("racket", "service", "get_racket_by_id")
	t.Parallel()
	t.WithNewStep("Success: correct id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.GetIncorrectID()
		racketTmp := utils.RacketObjectMother{}.DefaultRacket()

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetRacketByID", ctx, req).
			Return(racketTmp, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).GetRacketByID(ctx, req)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().NoError(err)
	})
}

// GetAllRackets
func (s *RacketServiceSuite) TestRacketServiceGetAllRackets1(t provider.T) {
	t.Title("[GetAllRackets] no rackets")
	t.Tags("racket", "service", "get_all_rackets")
	t.Parallel()
	t.WithNewStep("Incorrect: no rackets", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.IncorrectFieldToSort()

		racketMockRepo := mocks.NewIRacketRepository(t)
		racketMockRepo.
			On("GetAllRackets", ctx, req).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		rackets, err := service.NewRacketService(utils.NewMockLogger(), racketMockRepo).GetAllRackets(ctx, req)

		sCtx.Assert().Nil(rackets)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *RacketServiceSuite) TestRacketServiceGetAllRackets2(t provider.T) {
	t.Title("[GetAllRackets] get rackets")
	t.Tags("racket", "service", "get_all_rackets")
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
