package utils

import (
	"context"
	"time"

	"src/internal/model"
	"src/internal/repository"
	mypostgres "src/internal/repository/postgres"
	"src/pkg/storage/postgres"
	"src/pkg/utils"
)

const connURL = "postgresql://postgres:admin@localhost:5432/tests"

var ids map[string]int

func NewTestStorage() (*postgres.Postgres, map[string]int) {

	conn, err := postgres.New(connURL)
	if err != nil {
		panic(err)
	}

	ids = map[string]int{}
	ids["userID"] = initUserRepository(mypostgres.NewUserRepository(conn))
	ids["racketID"] = initRacketRepository(mypostgres.NewRacketRepository(conn))
	ids["orderID"] = initOrderRepository(mypostgres.NewOrderRepository(conn))
	return conn, ids
}

func DropTestStorage(testDB *postgres.Postgres) {
	defer testDB.Close()

	err := mypostgres.NewUserRepository(testDB).Delete(context.TODO(), ids["userID"])
	if err != nil {
		panic(err)
	}

	err = mypostgres.NewCartRepository(testDB).Delete(context.TODO(), ids["cartID"])
	if err != nil {
		panic(err)
	}

	err = mypostgres.NewRacketRepository(testDB).Delete(context.TODO(), ids["racketID"])
	if err != nil {
		panic(err)
	}
}

func initUserRepository(repo repository.IUserRepository) int {

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivan@mail.ru",
		Role:     model.UserRoleCustomer,
		Password: utils.HashAndSalt([]byte("ivan")),
	}

	err := repo.Create(context.TODO(), user)

	if err != nil {
		panic(err)
	}

	return user.ID
}

func initRacketRepository(repo repository.IRacketRepository) int {

	racket := &model.Racket{
		Brand:     "head",
		Weight:    100,
		Balance:   100,
		HeadSize:  100,
		Avaliable: true,
		Quantity:  100,
		Price:     100,
	}

	err := repo.Create(context.TODO(), racket)

	if err != nil {
		panic(err)
	}

	return racket.ID
}

func initOrderRepository(repo repository.IOrderRepository) int {

	tm, _ := time.Parse(time.RFC3339, "2006-01-02")
	order := &model.Order{
		UserID:        ids["userID"],
		CreationDate:  tm,
		Address:       "Moscow",
		RecepientName: "Stepan Postnov",
		Status:        model.OrderStatusInProgress,
		Lines: []*model.OrderLine{
			{
				RacketID: ids["racketID"],
				Quantity: 1,
			},
		},
	}

	err := repo.Create(context.TODO(), order)

	if err != nil {
		panic(err)
	}

	return order.ID
}
