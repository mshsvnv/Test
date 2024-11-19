package e2e_test

import (
	"net/http"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gavv/httpexpect/v2"
)

var (
	expectReset *httpexpect.Expect
)

func TestReset(t *testing.T) {

	clnt := &http.Client{}
	expectReset = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8082/api/v2/auth",
		Client:   clnt,
		Reporter: httpexpect.NewRequireReporter(nil),
	})

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeResetScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/reset.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run reset feature tests")
	}
}

func Reset(ctx *godog.ScenarioContext) {

	var response *httpexpect.Response

	recepientEmail := "stepaha78@gmail.com"
	recepientPassword := "admin"

	ctx.When(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectReset.Request(method, endpoint).
			WithJSON(map[string]string{
				"email":        recepientEmail,
				"old_password": recepientPassword,
			}).
			Expect()
		return nil
	})

	ctx.Then(`^the response code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"msg": "OTP code to \"Login\" was sent to your email",
		})
		return nil
	})
}

func VerifyOTPReset(ctx *godog.ScenarioContext) {

	var response *httpexpect.Response

	recepientEmail := "stepaha78@gmail.com"
	// recepientCode := "123456"
	// recepientPasswordNew := "admin"

	ctx.When(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectReset.Request(method, endpoint).
			WithJSON(map[string]string{
				"new_password": recepientEmail,
				"code":         recepientEmail,
				"email":        recepientEmail,
			}).
			Expect()
		return nil
	})

	ctx.Then(`^the response code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"access_token": "your_access_token",
		})
		return nil
	})
}

func InitializeResetScenario(ctx *godog.ScenarioContext) {
	Reset(ctx)
	VerifyOTPReset(ctx)
}
