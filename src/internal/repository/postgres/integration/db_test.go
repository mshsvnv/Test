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
	db, ctr, ids := utils.NewTestStorage()
	defer utils.DropTestStorage(db, ctr)

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
		// &FeedbackRepoSuite{},
		// &CartRepoSuite{},
		// &OrderRepoSuite{},
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
