package mypostgres

// import (
// 	"context"
// 	"os"
// 	"src/internal/model"
// 	"src/internal/repository"
// 	"src/pkg/storage/postgres"
// 	"src/pkg/utils"
// 	"testing"
// )

// const connURL = "postgresql://postgres:admin@localhost:5432/Shop"

// var testDB *postgres.Postgres
// var ids map[string]int
// var emails map[string]string

// func TestMain(m *testing.M) {

// 	testDB = NewTestStorage()

// 	code := m.Run()
// 	// DropTestStorage(testDB, ids)
// 	testDB.Close()

// 	os.Exit(code)
// }

// func NewTestStorage() *postgres.Postgres {

// 	conn, err := postgres.New(connURL)

// 	if err != nil {
// 		panic(err)
// 	}

// 	ids = map[string]int{}
// 	emails = map[string]string{}

// 	ids["userID"] = initTestUserStorage(NewUserRepository(conn))
// 	ids["racketID"] = initTestRacketStorage(NewRacketRepository(conn))

// 	return conn
// }

// // func DropTestStorage(testDB *postgres.Postgres, ids map[string]int) {

// // 	err := NewUserRepository(testDB).
// // 		Delete(
// // 			context.TODO(),
// // 			emails["userEmail"])

// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	err = NewRacketRepository(testDB).
// // 		Delete(
// // 			context.TODO(),
// // 			ids["racketID"])

// // 	if err != nil {
// // 		panic(err)
// // 	}
// // }

// func initTestUserStorage(repo repository.IUserRepository) int {

// 	emails["userEmail"] = "mshsvnv@mail.ru"

// 	user := &model.User{
// 		Email:    emails["userEmail"],
// 		Password: utils.HashAndSalt([]byte("123")),
// 	}

// 	err := repo.Create(context.TODO(), user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return user.ID
// }

// func initTestRacketStorage(repo repository.IRacketRepository) int {

// 	racket := &model.Racket{
// 		Quantity:  100,
// 		Avaliable: true,
// 	}

// 	err := repo.Create(context.TODO(), racket)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return racket.ID
// }
