package e2e_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gavv/httpexpect/v2"
	"github.com/joho/godotenv"
)

var (
	expectLogin *httpexpect.Expect
)

func TestLogin(t *testing.T) {

	clnt := &http.Client{}
	expectLogin = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8111/api/v2/auth",
		Client:   clnt,
		Reporter: httpexpect.NewRequireReporter(nil),
	})

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeLoginScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/login.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run login feature tests")
	}
}

func Login(ctx *godog.ScenarioContext) {

	godotenv.Load()

	var response *httpexpect.Response
	recepientEmail := os.Getenv("RECEPIENT_EMAIL_ADDRESS")

	// recepientEmail := "stepaha78@gmail.com"

	ctx.When(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectLogin.Request(method, endpoint).
			WithJSON(map[string]string{
				"email":    recepientEmail,
				"password": "admin",
			}).
			Expect()
		return nil
	})

	ctx.Then(`^the response code on /login should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /login should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"msg": "OTP code to \"Login\" was sent to your email",
		})
		return nil
	})

	ctx.Step(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectLogin.Request(method, endpoint).
			WithJSON(map[string]string{
				"email": recepientEmail,
				"code":  "123456",
			}).
			Expect()
		return nil
	})

	ctx.Then(`^the response code on /login/verify should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /login/verify should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"access_token": "your_access_token",
		})
		return nil
	})
}

func InitializeLoginScenario(ctx *godog.ScenarioContext) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return
	// }

	Login(ctx)
}
