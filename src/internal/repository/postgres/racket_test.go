package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/repository/utils"
)

type RacketRepoSuite struct {
	suite.Suite

	racketMockRepo mocks.IRacketRepository
}

func (r *RacketRepoSuite) BeforeAll(t provider.T) {
	t.Title("Init racket mock repo")
	r.racketMockRepo = *mocks.NewIRacketRepository(t)
	t.Tags("fixture", "racket")
}

func (r *RacketRepoSuite) TestRacketRepoCreate(t provider.T) {
	t.Title("[Create] Create racket")
	t.Tags("racket repository", "postgres", "")
	t.Parallel()
	t.WithNewStep("Create racket", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.DefaultRacket()

		r.racketMockRepo.
			On("Create", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := r.racketMockRepo.Create(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoUpdate(t provider.T) {
	t.Title("[Update] Update racket")
	t.Tags("racket repository", "postgres", "")
	t.Parallel()
	t.WithNewStep("Update racket", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.DefaultRacket()

		r.racketMockRepo.
			On("Update", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := r.racketMockRepo.Update(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoDelete(t provider.T) {
	t.Title("[Delete] Delete racket")
	t.Tags("racket repository", "postgres", "")
	t.Parallel()
	t.WithNewStep("Delete racket", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.DefaultRacket().ID

		r.racketMockRepo.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := r.racketMockRepo.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoGetRacketByID(t provider.T) {
	t.Title("[GetRacketByID] Get racket by ID")
	t.Tags("racket repository", "postgres", "")
	t.Parallel()
	t.WithNewStep("Get racket by ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.DefaultRacket().ID

		expRacket := &model.Racket{
			ID:        1,
			Price:     100,
			Quantity:  100,
			Avaliable: true,
		}

		r.racketMockRepo.
			On("GetRacketByID", ctx, request).
			Return(expRacket, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		racket, err := r.racketMockRepo.GetRacketByID(ctx, request)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(racket, expRacket)
	})
}

func (r *RacketRepoSuite) TestRacketRepoGetAllRackets(t provider.T) {
	t.Title("[GetAllRackets] Get all rackets")
	t.Tags("racket repository", "postgres", "")
	t.Parallel()
	t.WithNewStep("Get all rackets", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.SortByPriceReq()

		expRackets := []*model.Racket{
			{
				ID:        1,
				Price:     100,
				Quantity:  100,
				Avaliable: true,
			},
		}

		r.racketMockRepo.
			On("GetAllRackets", ctx, request).
			Return(expRackets, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		rackets, err := r.racketMockRepo.GetAllRackets(ctx, request)

		sCtx.Assert().NotEmpty(rackets)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(rackets, expRackets)
	})
}
