package mypostgres_test

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	wg := &sync.WaitGroup{}
	suites := []runner.TestSuite{
		&UserRepoSuite{},
		&RacketRepoSuite{},
		&FeedbackRepoSuite{},
		&CartRepoSuite{},
		&OrderRepoSuite{},
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
