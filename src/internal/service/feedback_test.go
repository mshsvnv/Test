package service_test

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
)

type FeedbackServiceSuite struct {
	suite.Suite
}

// CreateFeedback
func (s *FeedbackServiceSuite) TestFeedbackServiceCreateFeedback1(t provider.T) {
	t.Title("[CreateFeedback] Incorrect")
	t.Tags("feedback", "service", "create_feedback")
	t.Parallel()
	t.WithNewStep("Incorrect: repository error", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2006-01-02")
		req := utils.FeedbackBuilder{}.
			WithRacketID(1).
			WithUserID(1).
			WithRating(5).
			WithDate(tm).
			ToCreateDTO()

		expFeedback := &model.Feedback{
			RacketID: 1,
			UserID:   1,
			Rating:   5,
			Date:     tm,
		}
		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("Create", ctx, expFeedback).
			Return(fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		feedback, err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).CreateFeedback(ctx, req)

		sCtx.Assert().Nil(feedback)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *FeedbackServiceSuite) TestFeedbackServiceCreateFeedback2(t provider.T) {
	t.Title("[CreateFeedback] Correct")
	t.Tags("feedback", "service", "create_feedback")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2006-01-02")
		req := utils.FeedbackBuilder{}.
			WithRacketID(1).
			WithUserID(1).
			WithRating(5).
			WithDate(tm).
			ToCreateDTO()

		expFeedback := &model.Feedback{
			RacketID: 1,
			UserID:   1,
			Rating:   5,
			Date:     tm,
		}

		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("Create", ctx, expFeedback).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		feedback, err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).CreateFeedback(ctx, req)

		sCtx.Assert().NotEmpty(feedback)
		sCtx.Assert().Nil(err)
	})
}

// DeleteFeedback
func (s *FeedbackServiceSuite) TestFeedbackServiceDeleteFeedback1(t provider.T) {
	t.Title("[DeleteFeedback] not existed feedback")
	t.Tags("feedback", "service", "delete_feedback")
	t.Parallel()
	t.WithNewStep("Incorrect: not existed feedback", func(sCtx provider.StepCtx) {

		ctx := context.TODO()

		req := utils.FeedbackBuilder{}.
			WithRacketID(1).
			WithUserID(1)
		reqGet := req.ToGetDTO()
		reqDelete := req.ToDeleteDTO()

		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("GetFeedback", ctx, reqGet).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", reqDelete)

		err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).DeleteFeedback(ctx, reqDelete)

		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *FeedbackServiceSuite) TestFeedbackServiceDeleteFeedback2(t provider.T) {
	t.Title("[DeleteFeedback] repository error")
	t.Tags("feedback", "service", "delete_feedback")
	t.Parallel()
	t.WithNewStep("Incorrect: repository error", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		req := utils.FeedbackBuilder{}.
			WithRacketID(1).
			WithUserID(1)
		reqGet := req.ToGetDTO()
		reqDelete := req.ToDeleteDTO()

		feedback := &model.Feedback{
			RacketID: 1,
			UserID:   1,
			Rating:   5,
			Date:     tm,
		}

		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("GetFeedback", ctx, reqGet).
			Return(feedback, nil).
			Once()

		feedbackMockRepo.
			On("Delete", ctx, reqDelete).
			Return(fmt.Errorf("delete feedback fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", reqDelete)

		err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).DeleteFeedback(ctx, reqDelete)

		sCtx.Assert().Error(err)
	})
}

func (s *FeedbackServiceSuite) TestFeedbackServiceDeleteFeedback3(t provider.T) {
	t.Title("[DeleteFeedback] success")
	t.Tags("feedback", "service", "delete_feedback")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		tm, _ := time.Parse(time.RFC3339, "2006-01-02")

		req := utils.FeedbackBuilder{}.
			WithRacketID(1).
			WithUserID(1)
		reqGet := req.ToGetDTO()
		reqDelete := req.ToDeleteDTO()

		feedback := &model.Feedback{
			RacketID: 1,
			UserID:   1,
			Rating:   5,
			Date:     tm,
		}

		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("GetFeedback", ctx, reqGet).
			Return(feedback, nil).
			Once()

		feedbackMockRepo.
			On("Delete", ctx, reqDelete).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", reqDelete)

		err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).DeleteFeedback(ctx, reqDelete)

		sCtx.Assert().Nil(err)
	})
}

// GetFeedbacksByRacketID
func (s *FeedbackServiceSuite) TestFeedbackServiceGetFeedbacksByRacketID1(t provider.T) {
	t.Title("[GetFeedbacksByRacketID] Incorrect racket ID")
	t.Tags("feedbacks", "get_feedbacks_by_racket_id")
	t.Parallel()
	t.WithNewStep("Incorrect: not existed racket id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := 1
		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("GetFeedbacksByRacketID", ctx, req).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		feedbacks, err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).GetFeedbacksByRacketID(ctx, req)

		sCtx.Assert().Nil(feedbacks)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *FeedbackServiceSuite) TestFeedbackServiceGetFeedbacksByRacketID2(t provider.T) {
	t.Title("[GetFeedbacksByRacketID] success")
	t.Tags("feedbacks", "get_feedbacks_by_racket_id")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := 1
		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		expFeedbacks := []*model.Feedback{
			{
				RacketID: req,
				UserID:   1,
				Rating:   1,
			},
		}

		feedbackMockRepo.
			On("GetFeedbacksByRacketID", ctx, req).
			Return(expFeedbacks, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		feedbacks, err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).GetFeedbacksByRacketID(ctx, req)

		sCtx.Assert().NotEmpty(feedbacks)
		sCtx.Assert().Nil(err)
	})
}

// GetFeedbacksByUserID
func (s *FeedbackServiceSuite) TestFeedbackServiceGetFeedbacksByUserID1(t provider.T) {
	t.Title("[GetFeedbacksByUserID] Incorrect user ID")
	t.Tags("feedbacks", "get_feedbacks_by_user_id")
	t.Parallel()
	t.WithNewStep("Incorrect: not existed user id", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := 1
		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		feedbackMockRepo.
			On("GetFeedbacksByUserID", ctx, req).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		feedbacks, err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).GetFeedbacksByUserID(ctx, req)

		sCtx.Assert().Nil(feedbacks)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *FeedbackServiceSuite) TestFeedbackServiceGetFeedbacksByUserID2(t provider.T) {
	t.Title("[GetFeedbacksByUserID] Incorrect user ID")
	t.Tags("user", "get_feedbacks_by_user_id")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := 1
		feedbackMockRepo := mocks.NewIFeedbackRepository(t)

		expFeedbacks := []*model.Feedback{
			{
				UserID:   req,
				RacketID: 1,
				Rating:   1,
			},
		}

		feedbackMockRepo.
			On("GetFeedbacksByUserID", ctx, req).
			Return(expFeedbacks, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		feedbacks, err := service.NewFeedbackService(utils.NewMockLogger(), feedbackMockRepo).GetFeedbacksByUserID(ctx, req)

		sCtx.Assert().NotEmpty(feedbacks)
		sCtx.Assert().Nil(err)
	})
}
