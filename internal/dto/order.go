package dto

import (
	"time"

	"src/internal/model"
	"src/pkg/storage/postgres"
)

type UpdateOrderReq struct {
	OrderID int               `json:"order_id"`
	Status  model.OrderStatus `json:"status"`
}

type PlaceOrderReq struct {
	UserID        int       `json:"user_id"`
	DeliveryDate  time.Time `json:"delivery_date" format:"2006-01-02T15:04:05"`
	Address       string    `json:"address"`
	RecepientName string    `json:"recepient_name"`
}

type ListOrdersReq struct {
	Pagination *postgres.Pagination
}
