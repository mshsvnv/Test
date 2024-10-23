//go:build e2e

package e2e_test

import (
	"net/http"
	"sync"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/e2e"
	"src/internal/dto"
	"src/internal/model"
)

type E2ESuite struct {
	suite.Suite

	e        httpexpect.Expect
	racketID int
}

type AddRacketCartReq struct {
	Quantity int `json:"quantity"`
	RacketID int `json:"racket_id"`
}

type CartRes struct {
	Cart *model.Cart `json:"cart"`
}

func (s *E2ESuite) BeforeAll(t provider.T) {
	s.e = *httpexpect.WithConfig(httpexpect.Config{
		Client:   &http.Client{},
		BaseURL:  "http://localhost:8044/api/v2",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func (s *E2ESuite) TestE2E(t provider.T) {
	t.Title("[E2E] E2E Test")
	t.Tags("e2e")
	t.Parallel()
	t.WithNewStep("E2E Test", func(sCtx provider.StepCtx) {

		registerReq := &dto.RegisterReq{
			Name:     "Klim",
			Surname:  "Klimov",
			Email:    "klim@mail.ru",
			Password: "klim",
		}

		accessToken := s.e.POST("/auth/register").
			WithJSON(registerReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			ContainsKey("access_token").
			Value("access_token").Raw().(string)

		req := &AddRacketCartReq{
			RacketID: s.racketID,
			Quantity: 1,
		}

		var cart *CartRes

		s.e.POST("/cart").
			WithHeader("Authorization", "Bearer "+accessToken).
			WithJSON(req).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Decode(&cart)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Equal(cart.Cart.Quantity, 1)
		sCtx.Assert().Equal(cart.Cart.TotalPrice, float32(300))

		s.e.GET("/cart").
			WithHeader("Authorization", "Bearer "+accessToken).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Decode(&cart)

		sCtx.Assert().NotEmpty(cart)
		sCtx.Assert().Equal(cart.Cart.Quantity, 1)
		sCtx.Assert().Equal(cart.Cart.TotalPrice, float32(300))
	})
}

func TestRunner(t *testing.T) {
	db, ids := e2e.NewTestStorage()
	t.Cleanup(func() {
		e2e.DropTestStorage(db)
	})

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&E2ESuite{
			racketID: ids["racketID"],
		},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func() {
			suite.RunSuite(t, s)
			wg.Done()
		}()
	}

	wg.Wait()
}
