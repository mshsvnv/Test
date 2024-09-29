package mypostgres_test

import (
	"context"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/repository/utils"
)

type FeedbackRepoSuite struct {
	suite.Suite

	feedbackMockRepo mocks.IFeedbackRepository
}

func (r *FeedbackRepoSuite) BeforeAll(t provider.T) {
	t.Title("Init feedback mock repo")
	r.feedbackMockRepo = *mocks.NewIFeedbackRepository(t)
	t.Tags("fixture", "feedback")
}

func (f *FeedbackRepoSuite) TestFeedbackRepoCreate(t provider.T) {
	t.Title("[Create] Create feedback")
	t.Tags("feedback repository", "postgres")
	t.Parallel()
	t.WithNewStep("Create feedback", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		tm, _ := time.Parse(time.RFC3339, "2006-01-02")
		request := utils.FeedbackBuilder{
			RacketID: 1,
			UserID:   1,
			Rating:   5,
			Feedback: "cool",
			Date:     tm,
		}.ToModel()

		f.feedbackMockRepo.
			On("Create", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := f.feedbackMockRepo.Create(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (f *FeedbackRepoSuite) TestFeedbackRepoUpdate(t provider.T) {
	t.Title("[Update] Update feedback")
	t.Tags("feedback repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update feedback", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		tm, _ := time.Parse(time.RFC3339, "2006-01-02")
		request := utils.FeedbackBuilder{
			RacketID: 1,
			UserID:   1,
			Rating:   5,
			Feedback: "cool",
			Date:     tm,
		}.ToModel()

		f.feedbackMockRepo.
			On("Update", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := f.feedbackMockRepo.Update(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (f *FeedbackRepoSuite) TestFeedbackRepoDelete(t provider.T) {
	t.Title("[Delete] Delete feedback")
	t.Tags("feedback repository", "postgres")
	t.Parallel()
	t.WithNewStep("Delete feedback", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FeedbackBuilder{
			RacketID: 1,
			UserID:   1,
		}.ToDeleteDTO()

		f.feedbackMockRepo.
			On("Delete", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := f.feedbackMockRepo.Delete(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (f *FeedbackRepoSuite) TestFeedbackRepoGetFeedback(t provider.T) {
	t.Title("[GetFeedback] Get feedback")
	t.Tags("feedback repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get feedback", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FeedbackBuilder{
			RacketID: 1,
			UserID:   1,
		}.ToGetDTO()

		expFeedback := &model.Feedback{
			RacketID: request.RacketID,
			UserID:   request.UserID,
			Rating:   5,
		}

		f.feedbackMockRepo.
			On("GetFeedback", ctx, request).
			Return(expFeedback, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		feedback, err := f.feedbackMockRepo.GetFeedback(ctx, request)

		sCtx.Assert().NotEmpty(feedback)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expFeedback, feedback)
	})
}

func (f *FeedbackRepoSuite) TestFeedbackRepoGetFeedbacksByRacketID(t provider.T) {
	t.Title("[GetFeedbacksByRacketID] Get feedbacks by racket id")
	t.Tags("feedback repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get feedbacks by racket id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FeedbackBuilder{
			RacketID: 1,
			UserID:   1,
		}

		expFeedbacks := []*model.Feedback{
			{
				RacketID: request.RacketID,
				UserID:   request.UserID,
				Rating:   5,
			},
		}

		f.feedbackMockRepo.
			On("GetFeedbacksByRacketID", ctx, request.RacketID).
			Return(expFeedbacks, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request.RacketID)

		feedbacks, err := f.feedbackMockRepo.GetFeedbacksByRacketID(ctx, request.RacketID)

		sCtx.Assert().NotEmpty(feedbacks)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expFeedbacks, feedbacks)
	})
}

func (f *FeedbackRepoSuite) TestFeedbackRepoGetFeedbacksByUserID(t provider.T) {
	t.Title("[GetFeedbacksByUserID] Get feedback feedback")
	t.Tags("feedback repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get feedback feedback", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.FeedbackBuilder{
			RacketID: 1,
			UserID:   1,
		}

		expFeedbacks := []*model.Feedback{
			{
				RacketID: request.RacketID,
				UserID:   request.UserID,
				Rating:   5,
			},
		}

		f.feedbackMockRepo.
			On("GetFeedbacksByUserID", ctx, request.RacketID).
			Return(expFeedbacks, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request.RacketID)

		feedbacks, err := f.feedbackMockRepo.GetFeedbacksByUserID(ctx, request.RacketID)

		sCtx.Assert().NotEmpty(feedbacks)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expFeedbacks, feedbacks)
	})
}
