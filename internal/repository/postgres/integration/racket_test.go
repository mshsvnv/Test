//go:build integration

package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	repo "src/internal/repository"
	"src/internal/repository/utils"
)

type RacketRepoSuite struct {
	suite.Suite

	racketRepo repo.IRacketRepository
	racketID   int
}

func (r *RacketRepoSuite) TestRacketRepoCreate(t provider.T) {
	t.Title("[Create] Create racket")
	t.Tags("racket", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Create racket", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.DefaultRacket()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := r.racketRepo.Create(ctx, request)
		sCtx.Assert().NoError(err)

		err = r.racketRepo.Delete(ctx, request.ID)
		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoUpdate(t provider.T) {
	t.Title("[Update] Update racket")
	t.Tags("racket", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update racket", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request, err := r.racketRepo.GetRacketByID(ctx, r.racketID)

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = r.racketRepo.Update(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoDelete(t provider.T) {
	t.Title("[Delete] Delete racket")
	t.Tags("racket", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Delete racket", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.DefaultRacket()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := r.racketRepo.Create(ctx, request)
		sCtx.Assert().NoError(err)

		err = r.racketRepo.Delete(ctx, request.ID)
		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoGetRacketByID(t provider.T) {
	t.Title("[GetRacketByID] Get racket by ID")
	t.Tags("racket", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get racket by ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := r.racketID

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		racket, err := r.racketRepo.GetRacketByID(ctx, request)

		sCtx.Assert().NotEmpty(racket)
		sCtx.Assert().NoError(err)
	})
}

func (r *RacketRepoSuite) TestRacketRepoGetAllRackets(t provider.T) {
	t.Title("[GetAllRackets] Get all rackets")
	t.Tags("racket", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get all rackets", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.RacketObjectMother{}.SortByPriceReq()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		rackets, err := r.racketRepo.GetAllRackets(ctx, request)

		sCtx.Assert().NotEmpty(rackets)
		sCtx.Assert().NoError(err)
	})
}
