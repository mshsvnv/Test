package utils

import (
	"src/internal/dto"
	"src/internal/model"
	"src/pkg/storage/postgres"
)

type AuthObjectMother struct {
}

func (o AuthObjectMother) DefaultUserReq() *dto.RegisterReq {
	return &dto.RegisterReq{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivan@mail.ru",
		Password: "ivan",
	}
}

func (o AuthObjectMother) RegisterNewUserReq() *dto.RegisterReq {
	return &dto.RegisterReq{
		Name:     "Peter",
		Surname:  "Petrov",
		Email:    "peter@mail.ru",
		Password: "peter",
	}
}

func (o AuthObjectMother) UnRegisterUserReq() *dto.LoginReq {
	return &dto.LoginReq{
		Email:    "vasya@mail.ru",
		Password: "vasya",
	}
}

func (o AuthObjectMother) IncorrectPasswordReq() *dto.LoginReq {
	return &dto.LoginReq{
		Email:    "ivan@mail.ru",
		Password: "peter",
	}
}

func (o AuthObjectMother) CorrectPasswordReq() *dto.LoginReq {
	return &dto.LoginReq{
		Email:    "ivan@mail.ru",
		Password: "ivan",
	}
}

type UserObjectMother struct {
}

func (o UserObjectMother) IncorrectID() int {
	return 0
}

func (o UserObjectMother) CorrectID() int {
	return 1 //ids["userID"]
}

func (o UserObjectMother) CorrectEmail() string {
	return "pstpn@mail.ru"
}

func (o UserObjectMother) DefaultCustomer() *model.User {
	return &model.User{
		ID:      1, //ids["userID"],
		Name:    "Ivan",
		Surname: "Ivanov",
		Email:   "ivan@mail.ru",
		Role:    model.UserRoleCustomer,
	}
}

func (o UserObjectMother) DefaultAdmin() *model.User {
	return &model.User{
		ID:      1, //ids["userID"],
		Name:    "Ivan",
		Surname: "Ivanov",
		Email:   "ivan@mail.ru",
		Role:    model.UserRoleAdmin,
	}
}

func (o UserObjectMother) DefaultUsers() []*model.User {
	return []*model.User{
		{
			Name:    "Ivan",
			Surname: "Ivanov",
			Email:   "ivan@mail.ru",
		},
		{
			Name:    "Peter",
			Surname: "Petrov",
			Email:   "peter@mail.ru",
		},
	}
}

func (o UserObjectMother) IncorrectUserIDToUpdate() *dto.UpdateRoleReq {
	return &dto.UpdateRoleReq{
		ID:   0,
		Role: model.UserRoleAdmin,
	}
}

func (o UserObjectMother) CorrectUserToUpdate() *dto.UpdateRoleReq {
	return &dto.UpdateRoleReq{
		ID:   1, //ids["userID"],
		Role: model.UserRoleAdmin,
	}
}

type RacketObjectMother struct {
}

func (r RacketObjectMother) DefaultRacket() *model.Racket {
	return &model.Racket{
		ID:        1,
		Price:     100,
		Quantity:  100,
		Avaliable: true,
	}
}

func (r RacketObjectMother) IncorrectCount() *dto.CreateRacketReq {
	return &dto.CreateRacketReq{
		Quantity: -1,
	}
}

func (r RacketObjectMother) CorrectCount() *dto.CreateRacketReq {
	return &dto.CreateRacketReq{
		Quantity: 10,
	}
}

func (r RacketObjectMother) UpdateIncorrectID() *dto.UpdateRacketReq {
	return &dto.UpdateRacketReq{
		ID:       0,
		Quantity: 1,
	}
}

func (r RacketObjectMother) UpdateCorrectID() *dto.UpdateRacketReq {
	return &dto.UpdateRacketReq{
		ID:       1, //ids["racketID"],
		Quantity: 100,
	}
}

func (r RacketObjectMother) GetIncorrectID() int {
	return 0
}

func (r RacketObjectMother) GetCorrectID() int {
	return 1 //ids["racketID"]
}

func (r RacketObjectMother) IncorrectFieldToSort() *dto.ListRacketsReq {
	return &dto.ListRacketsReq{
		Pagination: &postgres.Pagination{
			Filter: postgres.FilterOptions{
				Column: "",
			},
			Sort: postgres.SortOptions{
				Direction: postgres.ASC,
				Columns:   []string{""},
			},
		},
	}
}

func (r RacketObjectMother) SortByPriceReq() *dto.ListRacketsReq {
	return &dto.ListRacketsReq{
		Pagination: &postgres.Pagination{
			Sort: postgres.SortOptions{
				Direction: postgres.ASC,
				Columns:   []string{"price"},
			},
		},
	}
}

type CartObjectMother struct {
	UserID   int
	RacketID int
	Quantity int
}

func (c CartObjectMother) GetCartByID() int {
	return c.UserID
}

func (c CartObjectMother) AddCartRacketReq() *dto.AddRacketCartReq {
	return &dto.AddRacketCartReq{
		UserID:   c.UserID,
		RacketID: c.RacketID,
		Quantity: 1,
	}
}

func (c CartObjectMother) RemoveRacketReq() *dto.RemoveRacketCartReq {
	return &dto.RemoveRacketCartReq{
		UserID:   c.UserID,
		RacketID: c.RacketID,
	}
}

func (c CartObjectMother) UpdatePlusRacketReq() *dto.UpdateRacketCartReq {
	return &dto.UpdateRacketCartReq{
		UserID:   c.UserID,
		RacketID: c.RacketID,
		Quantity: 1,
	}
}

func (c CartObjectMother) UpdateRacketMinusReq() *dto.UpdateRacketCartReq {
	return &dto.UpdateRacketCartReq{
		UserID:   c.UserID,
		RacketID: c.RacketID,
		Quantity: -1,
	}
}

func (c CartObjectMother) DefaultCart() *model.Cart {
	return &model.Cart{
		UserID:   c.UserID,
		Quantity: c.Quantity,
		Lines: []*model.CartLine{
			{
				RacketID: c.RacketID,
				Quantity: c.Quantity,
			},
		},
	}
}
