package service_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/service"
	"src/internal/service/utils"
)

type RacketSuite struct {
	suite.Suite

	racketService service.IRacketService
}

// CreateRacket
func (s *RacketSuite) TestRacketServiceCreateRacket1(t provider.T) {
	t.Title("[Integration CreateRacket] wrong amount")
	t.Tags("integration",  "racket", "service", "create_racket")
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

func (s *RacketSuite) TestRacketServiceCreateRacket2(t provider.T) {
	t.Title("[Integration CreateRacket] correct request")
	t.Tags("integration",  "racket", "service", "create_racket")
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
func (s *RacketSuite) TestRacketServiceUpdateRacket1(t provider.T) {
	t.Title("[Integration UpdateRacket] wrong ID")
	t.Tags("integration",  "racket", "service", "update_racket")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong ID", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.UpdateIncorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := s.racketService.UpdateRacket(ctx, req)

		sCtx.Assert().Error(err)
	})
}

func (s *RacketSuite) TestRacketServiceUpdateRacket2(t provider.T) {
	t.Title("[Integration UpdateRacket] correct request")
	t.Tags("integration",  "racket", "service", "update_racket")
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
func (s *RacketSuite) TestRacketServiceGetRacketByID1(t provider.T) {
	t.Title("[Integration GetRacketByID] wrong id")
	t.Tags("integration",  "racket", "service", "get_racket_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: wrong id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.GetIncorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := s.racketService.GetRacketByID(ctx, req)

		sCtx.Assert().Nil(racket)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *RacketSuite) TestRacketServiceGetRacketByID2(t provider.T) {
	t.Title("[Integration GetRacketByID] correct id")
	t.Tags("integration",  "racket", "service", "get_racket_by_id")
	t.Parallel()
	t.WithNewStep("Success: correct id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.GetCorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		racket, err := s.racketService.GetRacketByID(ctx, req)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().Nil(err)
	})
}

// GetAllRackets
func (s *RacketSuite) TestRacketServiceGetAllRackets1(t provider.T) {
	t.Title("[Integration GetAllRackets] no rackets")
	t.Tags("integration",  "racket", "service", "get_all_rackets")
	t.Parallel()
	t.WithNewStep("Incorrect: no rackets", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.RacketObjectMother{}.IncorrectFieldToSort()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		rackets, err := s.racketService.GetAllRackets(ctx, req)

		sCtx.Assert().Nil(rackets)
		sCtx.Assert().Error(err)
	})
}
