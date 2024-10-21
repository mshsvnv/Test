//go:build unit

package service_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

var signingKey = "racket_shop"
var accessTokenTTL time.Duration = time.Duration(12 * time.Hour.Hours())

func TestRunner(t *testing.T) {

	t.Parallel()

	wg := &sync.WaitGroup{}
	suites := []runner.TestSuite{
		&AuthServiceSuite{},
		&UserServiceSuite{},
		&RacketServiceSuite{},
		&FeedbackServiceSuite{},
		&CartServiceSuite{},
		&OrderServiceSuite{},
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
