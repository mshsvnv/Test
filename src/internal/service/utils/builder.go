package utils

import (
	"src/internal/dto"
	"src/internal/model"
	"time"
)

type FeedbackBuilder struct {
	RacketID int
	UserID   int
	Feedback string
	Rating   int
	Date     time.Time
}

func (f FeedbackBuilder) WithRacketID(racketID int) FeedbackBuilder {
	f.RacketID = racketID
	return f
}

func (f FeedbackBuilder) WithUserID(userID int) FeedbackBuilder {
	f.UserID = userID
	return f
}

func (f FeedbackBuilder) WithFeedback(feedback string) FeedbackBuilder {
	f.Feedback = feedback
	return f
}

func (f FeedbackBuilder) WithDate(date time.Time) FeedbackBuilder {
	f.Date = date
	return f
}

func (f FeedbackBuilder) WithRating(rating int) FeedbackBuilder {
	f.Rating = rating
	return f
}

func (f FeedbackBuilder) ToModel() *model.Feedback {
	return &model.Feedback{
		RacketID: f.RacketID,
		UserID:   f.UserID,
		Feedback: f.Feedback,
		Rating:   f.Rating,
		Date:     f.Date,
	}
}

func (f FeedbackBuilder) ToCreateDTO() *dto.CreateFeedbackReq {
	return &dto.CreateFeedbackReq{
		RacketID: f.RacketID,
		UserID:   f.UserID,
		Feedback: f.Feedback,
		Rating:   f.Rating,
	}
}

func (f FeedbackBuilder) ToGetDTO() *dto.GetFeedbackReq {
	return &dto.GetFeedbackReq{
		RacketID: f.RacketID,
		UserID:   f.UserID,
	}
}

func (f FeedbackBuilder) ToDeleteDTO() *dto.DeleteFeedbackReq {
	return &dto.DeleteFeedbackReq{
		RacketID: f.RacketID,
		UserID:   f.UserID,
	}
}
