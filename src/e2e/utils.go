package e2e

import (
	"context"

	"src/internal/model"
	"src/internal/repository"
	mypostgres "src/internal/repository/postgres"
	"src/pkg/storage/postgres"
)

const connString = "postgresql://postgres:admin@localhost:5434/tests"

var ids map[string]int

func NewTestStorage() (*postgres.Postgres, map[string]int) {
	conn, err := postgres.New(connString)
	if err != nil {
		panic(err)
	}

	ids = map[string]int{}
	ids["racketID"] = initRacketRepository(mypostgres.NewRacketRepository(conn))
	return conn, ids
}

func DropTestStorage(testDB *postgres.Postgres) {
	defer func() {
		testDB.Close()
	}()

	err := mypostgres.NewRacketRepository(testDB).Delete(context.TODO(), ids["racketID"])
	if err != nil {
		panic(err)
	}
}

func initRacketRepository(repo repository.IRacketRepository) int {

	racket := &model.Racket{
		Brand:     "babolat",
		Weight:    300,
		Balance:   300,
		HeadSize:  300,
		Avaliable: true,
		Quantity:  300,
		Price:     300,
	}

	err := repo.Create(context.TODO(), racket)

	if err != nil {
		panic(err)
	}

	return racket.ID
}
