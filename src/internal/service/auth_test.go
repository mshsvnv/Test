package service_test

import (
	"context"
	"fmt"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"
)

type AuthServiceSuite struct {
	suite.Suite
}

// func (s *AuthServiceSuite) TestAuthServiceRegister1(t provider.T) {
// 	t.Title("[Register] Correct")
// 	t.Tags("auth", "service", "register")
// 	t.Parallel()
// 	t.WithNewStep("Correct", func(sCtx provider.StepCtx) {

// 		ctx := context.TODO()
// 		req := utils.AuthObjectMother{}.DefaultUserReq()

// 		userMockRepo := mocks.NewIUserRepository(t)
// 		userMockRepo.
// 			On("GetUserByEmail", ctx, req.Email).
// 			Return(nil, nil).
// 			Once()

// 		user := utils.AuthObjectMother{}.DefaultUserModel()
// 		userMockRepo.
// 			On("Create", ctx, user).
// 			Return(nil).
// 			Once()

// 		sCtx.WithNewParameters("ctx", ctx, "request", req)

// 		token, err := service.NewAuthService(utils.NewMockLogger(), userMockRepo, signingKey, accessTokenTTL).Register(ctx, req)

// 		sCtx.Assert().NotEmpty(token)
// 		sCtx.Assert().NoError(err)
// 	})
// }

func (s *AuthServiceSuite) TestAuthServiceRegister2(t provider.T) {
	t.Title("[Register] User already exists")
	t.Tags("auth", "service", "register")
	t.Parallel()
	t.WithNewStep("Incorrect: user already exists", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.DefaultUserReq()
		user := utils.AuthObjectMother{}.DefaultUserModel()

		userMockRepo := mocks.NewIUserRepository(t)
		userMockRepo.
			On("GetUserByEmail", ctx, req.Email).
			Return(user, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := service.NewAuthService(utils.NewMockLogger(), userMockRepo, signingKey, accessTokenTTL).Register(ctx, req)

		sCtx.Assert().Empty(token)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "get user by email fail, error already exists")
	})
}

func (s *AuthServiceSuite) TestAuthServiceLogin1(t provider.T) {
	t.Title("[Login] Correct password")
	t.Tags("auth", "service", "login")
	t.Parallel()
	t.WithNewStep("Correct: correct password", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.CorrectPasswordReq()
		user := utils.AuthObjectMother{}.DefaultUserModel()

		userMockRepo := mocks.NewIUserRepository(t)
		userMockRepo.
			On("GetUserByEmail", ctx, req.Email).
			Return(user, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := service.NewAuthService(utils.NewMockLogger(), userMockRepo, signingKey, accessTokenTTL).Login(ctx, req)

		sCtx.Assert().NotEmpty(token)
		sCtx.Assert().NoError(err)
	})
}

func (s *AuthServiceSuite) TestAuthServiceLogin2(t provider.T) {
	t.Title("[Login] User doesn't exist")
	t.Tags("auth", "service", "login")
	t.Parallel()
	t.WithNewStep("Incorrect: user doesn't exist", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.UnRegisterUserReq()

		userMockRepo := mocks.NewIUserRepository(t)
		userMockRepo.
			On("GetUserByEmail", ctx, req.Email).
			Return(nil, fmt.Errorf("get user by email fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := service.NewAuthService(utils.NewMockLogger(), userMockRepo, signingKey, accessTokenTTL).Login(ctx, req)

		sCtx.Assert().Empty(token)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "get user by email fail")
	})
}

func (s *AuthServiceSuite) TestAuthServiceLogin3(t provider.T) {
	t.Title("[Login] Incorrect password")
	t.Tags("auth", "service", "login")
	t.Parallel()
	t.WithNewStep("Incorrect: incorrect password", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.AuthObjectMother{}.UnRegisterUserReq()
		user := utils.AuthObjectMother{}.DefaultUserModel()

		userMockRepo := mocks.NewIUserRepository(t)
		userMockRepo.
			On("GetUserByEmail", ctx, req.Email).
			Return(user, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		token, err := service.NewAuthService(utils.NewMockLogger(), userMockRepo, signingKey, accessTokenTTL).Login(ctx, req)

		sCtx.Assert().Empty(token)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "wrong password, error")
	})
}
