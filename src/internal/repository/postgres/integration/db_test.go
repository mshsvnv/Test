//go:build integration

package mypostgres_test

import (
	mypostgres "src/internal/repository/postgres"
	"src/internal/repository/postgres/utils"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestRunner(t *testing.T) {
	db, ctr, _ := utils.NewTestStorage()
	defer utils.DropTestStorage(db, ctr)

	t.Parallel()

	wg := &sync.WaitGroup{}
	suites := []runner.TestSuite{
		&UserRepoSuite{
			userRepo: mypostgres.NewUserRepository(db),
		},
		// &RacketRepoSuite{},
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
