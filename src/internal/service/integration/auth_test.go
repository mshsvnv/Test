//go:build integration

package service_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	repo "src/internal/repository"
	"src/internal/service"
	"src/internal/service/utils"
)

type AuthSuite struct {
	suite.Suite

	authService service.IAuthService
	userRepo    repo.IUserRepository
	userID      int
}

func (s *AuthSuite) TestAuthServiceRegister1(t provider.T) {
	t.Title("[Register] Successfully registered")
	t.Tags("integration", "auth", "service", "register")
	t.Parallel()
	t.WithNewStep("Correct: successfully registered", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.RegisterNewUserReq()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := s.authService.Register(ctx, req)

		sCtx.Assert().NotEmpty(token)
		sCtx.Assert().Nil(err)

		user, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
		err = s.userRepo.Delete(ctx, user.ID)
		sCtx.Assert().Nil(err)
	})
}

func (s *AuthSuite) TestAuthServiceRegister2(t provider.T) {
	t.Title("[Register] User already exists")
	t.Tags("integration", "auth", "service", "register")
	t.Parallel()
	t.WithNewStep("Incorrect: user already exists", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.DefaultUserReq()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := s.authService.Register(ctx, req)

		sCtx.Assert().Empty(token)
		sCtx.Assert().Error(err)
	})
}

func (s *AuthSuite) TestAuthServiceLogin1(t provider.T) {
	t.Title("[Login] Correct password")
	t.Tags("integration", "auth", "service", "login")
	t.Parallel()
	t.WithNewStep("Correct: correct password", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.CorrectPasswordReq()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := s.authService.Login(ctx, req)

		sCtx.Assert().NotEmpty(token)
		sCtx.Assert().Nil(err)
	})
}

func (s *AuthSuite) TestAuthServiceLogin2(t provider.T) {
	t.Title("[Login] User doesn't exist")
	t.Tags("integration", "auth", "service", "login")
	t.Parallel()
	t.WithNewStep("Incorrect: user doesn't exist", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.UnRegisterUserReq()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := s.authService.Login(ctx, req)

		sCtx.Assert().Empty(token)
		sCtx.Assert().Error(err)
	})
}

func (s *AuthSuite) TestAuthServiceLogin3(t provider.T) {
	t.Title("[Login] Incorrect password")
	t.Tags("integration", "auth", "service", "login")
	t.Parallel()
	t.WithNewStep("Incorrect: incorrect password", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.IncorrectPasswordReq()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := s.authService.Login(ctx, req)

		sCtx.Assert().Empty(token)
		sCtx.Assert().Error(err)
	})
}
