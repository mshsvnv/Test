//go:build integration

package mypostgres_test

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	mypostgres "src/internal/repository/postgres"
	"src/internal/repository/postgres/utils"
)

func TestRunner(t *testing.T) {
	db, ids := utils.NewTestStorage()
	defer utils.DropTestStorage(db)

	t.Parallel()

	wg := &sync.WaitGroup{}
	suites := []runner.TestSuite{
		&UserRepoSuite{
			userRepo: mypostgres.NewUserRepository(db),
			userID:   ids["userID"],
		},
		&RacketRepoSuite{
			racketRepo: mypostgres.NewRacketRepository(db),
			racketID:   ids["racketID"],
		},
		&CartRepoSuite{
			cartRepo: mypostgres.NewCartRepository(db),
			userRepo: mypostgres.NewUserRepository(db),
			cartID:   ids["cartID"],
			racketID: ids["racketID"],
		},
		&OrderRepoSuite{
			orderRepo: mypostgres.NewOrderRepository(db),
			orderID:   ids["orderID"],
			userID:    ids["userID"],
		},
	}
	wg.Add(len(suites))

	for _, s := range suites {
		go func() {
			suite.RunSuite(t, s)
			wg.Done()
		}()
	}

	wg.Wait()
}
