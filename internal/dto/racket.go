package dto

import "src/pkg/storage/postgres"

type CreateRacketReq struct {
	Brand     string  `json:"brand"`
	Weight    float32 `json:"weight"`
	Balance   float32 `json:"balance"`
	HeadSize  float32 `json:"head_size"`
	Avaliable bool    `json:"avaliable"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	Photo     []byte  `json:"photo"`
}

type UpdateRacketReq struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type ListRacketsReq struct {
	Pagination *postgres.Pagination
}
