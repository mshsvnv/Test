package utils

import (
	"context"

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
	ids["cartID"] = initCartRepository(mypostgres.NewCartRepository(conn))

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

	// err = mypostgres.NewOrderRepository(testDB).Delete(context.TODO(), ids["orderID"])
	// if err != nil {
	// 	panic(err)
	// }

	// err = mypostgres.NewFeedbackRepository(testDB).Delete(context.TODO(), &dto.RemoveFeedbackReq{
	// 	RacketID: ids["racketID"],
	// 	UserID:   ids["userID"],
	// })
	// if err != nil {
	// 	panic(err)
	// }
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

func initCartRepository(repo repository.ICartRepository) int {

	quantity := 1
	price := float32(1000)

	cart := &model.Cart{
		UserID:     ids["userID"],
		TotalPrice: price,
		Quantity:   quantity,
		Lines: []*model.CartLine{{
			RacketID: ids["racketID"],
			Quantity: quantity,
			Price:    price,
		}},
	}

	err := repo.Create(context.TODO(), cart)

	if err != nil {
		panic(err)
	}

	return ids["userID"]
}

// func initRacketRepository(repo mypostgres.UserRepository) int {

// 	user := &model.User{
// 		Name:     "Ivan",
// 		Surname:  "Ivanov",
// 		Email:    "ivan@mail.ru",
// 		Password: "ivan",
// 	}
// 	err := repo.Create(context.TODO(), user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return user.ID
// }

// func initOrderRepository(repo mypostgres.UserRepository) int {

// 	user := &model.User{
// 		Name:     "Ivan",
// 		Surname:  "Ivanov",
// 		Email:    "ivan@mail.ru",
// 		Password: "ivan",
// 	}
// 	err := repo.Create(context.TODO(), user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return user.ID
// }

// func initFeedbackRepository(repo mypostgres.UserRepository) int {

// 	user := &model.User{
// 		Name:     "Ivan",
// 		Surname:  "Ivanov",
// 		Email:    "ivan@mail.ru",
// 		Password: "ivan",
// 	}
// 	err := repo.Create(context.TODO(), user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return user.ID
// }
