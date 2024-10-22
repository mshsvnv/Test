package utils

import (
	"context"
	"os"
	"time"

	"github.com/testcontainers/testcontainers-go"
	container "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"src/internal/model"
	"src/internal/repository"
	mypostgres "src/internal/repository/postgres"
	"src/pkg/storage/postgres"
	"src/pkg/utils"
)

const (
	pgImage    = "docker.io/postgres:16-alpine"
	dbName     = "tests"
	dbUsername = "postgres"
	dbPassword = "admin"
)

var ids map[string]int

func NewTestStorage() (*postgres.Postgres, *container.PostgresContainer, map[string]int) {

	ctr, err := container.Run(
		context.TODO(),
		pgImage,
		container.WithInitScripts(os.Getenv("DB_INIT_PATH")),
		container.WithDatabase(dbName),
		container.WithUsername(dbUsername),
		container.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	connString, err := ctr.ConnectionString(context.TODO(), "sslmode=disable")
	if err != nil {
		panic(err)
	}

	conn, err := postgres.New(connString)
	if err != nil {
		panic(err)
	}

	ids = map[string]int{}
	ids["userID"] = initUserRepository(mypostgres.NewUserRepository(conn))
	ids["racketID"] = initRacketRepository(mypostgres.NewRacketRepository(conn))
	ids["orderID"] = initOrderRepository(mypostgres.NewOrderRepository(conn))
	ids["cartID"] = initCartRepository(mypostgres.NewCartRepository(conn))

	return conn, ctr, ids
}

func DropTestStorage(testDB *postgres.Postgres, ctr *container.PostgresContainer) {
	defer func() {
		testDB.Close()
		ctr.Terminate(context.TODO())
	}()

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

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:07:00Z")
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

func initCartRepository(repo repository.ICartRepository) int {

	cart := &model.Cart{
		UserID:     ids["userID"],
		TotalPrice: 100,
		Quantity:   1,
		Lines: []*model.CartLine{
			{
				RacketID: ids["racketID"],
				Quantity: 1,
			},
		},
	}

	err := repo.Create(context.TODO(), cart)

	if err != nil {
		panic(err)
	}

	return cart.UserID
}
