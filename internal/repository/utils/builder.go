package utils

import (
	"src/internal/dto"
	"src/internal/model"
	"src/pkg/storage/postgres"
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

type OrderBuilder struct {
	OrderID       int
	UserID        int
	CreationDate  time.Time
	DeliveryDate  time.Time
	Address       string
	RecepientName string
	Status        model.OrderStatus
	Lines         []*model.OrderLine
	TotalPrice    float32
}

func (f OrderBuilder) WithUserID() OrderBuilder {
	f.UserID = 1 // ids["userID"]
	return f
}

func (f OrderBuilder) WithOrderID() OrderBuilder {
	f.OrderID = 1// ids["orderID"]
	return f
}

func (f OrderBuilder) WithCreationDate(date time.Time) OrderBuilder {
	f.CreationDate = date
	return f
}

func (f OrderBuilder) WithDeliveryDate(date time.Time) OrderBuilder {
	f.DeliveryDate = date
	return f
}

func (f OrderBuilder) WithAddress(address string) OrderBuilder {
	f.Address = address
	return f
}

func (f OrderBuilder) WithRecepientName(recepient string) OrderBuilder {
	f.RecepientName = recepient
	return f
}

func (f OrderBuilder) WithStatus(status model.OrderStatus) OrderBuilder {
	f.Status = status
	return f
}

func (f OrderBuilder) WithTotalPrice(totalPrice float32) OrderBuilder {
	f.TotalPrice = totalPrice
	return f
}

func (f OrderBuilder) WithLines(lines []*model.OrderLine) OrderBuilder {
	f.Lines = lines
	return f
}

func (f OrderBuilder) ToUpdateDTO(status model.OrderStatus) *dto.UpdateOrderReq {
	return &dto.UpdateOrderReq{
		OrderID: f.OrderID,
		Status:  status,
	}
}

func (f OrderBuilder) ToPlaceOrderDTO() *dto.PlaceOrderReq {
	return &dto.PlaceOrderReq{
		UserID:        f.UserID,
		DeliveryDate:  f.DeliveryDate,
		Address:       f.Address,
		RecepientName: f.RecepientName,
	}
}

func (f OrderBuilder) ToModel(cart *model.Cart) *model.Order {

	var order model.Order

	for _, line := range cart.Lines {
		order.Lines = append(order.Lines,
			&model.OrderLine{
				RacketID: line.RacketID,
				Quantity: line.Quantity,
			})
	}

	order.UserID = f.UserID
	order.CreationDate = f.CreationDate
	order.DeliveryDate = f.DeliveryDate
	order.Address = f.Address
	order.RecepientName = f.RecepientName
	order.TotalPrice = cart.TotalPrice
	order.Status = f.Status

	return &order
}

func (f OrderBuilder) ToListAllOrders(columns []string) *dto.ListOrdersReq {
	return &dto.ListOrdersReq{
		Pagination: &postgres.Pagination{
			Sort: postgres.SortOptions{
				Direction: postgres.ASC,
				Columns:   columns,
			},
		},
	}
}
