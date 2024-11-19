package dto

import (
	"time"
)

type CreateFeedbackReq struct {
	RacketID int    `json:"racket_id"`
	UserID   int    `json:"user_id"`
	Feedback string `json:"feedback"`
	Rating   int    `json:"rating"`
}

type DeleteFeedbackReq struct {
	RacketID int `json:"racket_id"`
	UserID   int `json:"user_id"`
}

type GetFeedbackReq struct {
	RacketID int `json:"racket_id"`
	UserID   int `json:"user_id"`
}

type UpdateFeedbackReq struct {
	RacketID int       `json:"racket_id"`
	UserID   int       `json:"user_id"`
	Feedback string    `json:"feedback"`
	Date     time.Time `json:"date"`
	Rating   int       `json:"rating"`
}
