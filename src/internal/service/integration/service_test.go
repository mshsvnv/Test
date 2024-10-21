//go:build integration

package service_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	mypostgres "src/internal/repository/postgres"
	"src/internal/service"
	"src/internal/service/utils"
)

var signingKey = "racket_shop"
var accessTokenTTL time.Duration = time.Duration(12 * time.Hour.Hours())

func TestRunner(t *testing.T) {

	db, _ := utils.NewTestStorage()
	defer utils.DropTestStorage(db)

	t.Parallel()

	wg := &sync.WaitGroup{}
	suites := []runner.TestSuite{
		&AuthSuite{
			authService: service.NewAuthService(
				utils.NewMockLogger(),
				mypostgres.NewUserRepository(db),
				signingKey,
				accessTokenTTL,
			),
		},
		&UserSuite{
			userService: service.NewUserService(
				utils.NewMockLogger(),
				mypostgres.NewUserRepository(db),
			),
		},
		&RacketSuite{
			racketService: service.NewRacketService(
				utils.NewMockLogger(),
				mypostgres.NewRacketRepository(db),
			),
		},
		// &FeedbackSuite{
		// 	feedbackService: service.NewFeedbackService(
		// 		utils.NewMockLogger(),
		// 		mypostgres.NewFeedbackRepository(db),
		// 	),
		// },
		// &CartServiceSuite{},
		// &OrderServiceSuite{
		// 	orderService: service.NewOrderService(
		// 		utils.NewMockLogger(),
		// 		mypostgres.NewOrderRepository(db),
		// 		mypostgres.NewCartRepository(db),
		// 		mypostgres.NewRacketRepository(db),
		// 	),
		// },
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
