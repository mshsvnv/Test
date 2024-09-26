package service_test

import (
	"context"
	"fmt"
	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UserSuite struct {
	suite.Suite
}

// GetUserByID
func (s *UserSuite) TestGetUserByID1(t provider.T) {
	t.Title("[GetUserByID] Incorrect ID")
	t.Tags("user", "get_user_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: not existed user", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.IncorrectID()
		userMockRepo := mocks.NewIUserRepository(t)

		userMockRepo.
			On("GetUserByID", ctx, req).
			Return(nil, fmt.Errorf("incorrect user ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).GetUserByID(ctx, req)

		sCtx.Assert().Nil(user)
		sCtx.Assert().Error(err)
	})
}

func (s *UserSuite) TestGetUserByID2(t provider.T) {
	t.Title("[GetUserByID] Correct ID")
	t.Tags("user", "get_user_by_id")
	t.Parallel()
	t.WithNewStep("Correct: existed user", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.CorrectID()
		userMockRepo := mocks.NewIUserRepository(t)

		user := &model.User{
			Name:    "Ivan",
			Surname: "Ivanov",
			Email:   "ivan@mail.ru",
			Role:    model.UserRoleCustomer,
		}

		userMockRepo.
			On("GetUserByID", ctx, req).
			Return(user, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).GetUserByID(ctx, req)

		sCtx.Assert().NotEmpty(user)
		sCtx.Assert().Nil(err)
	})
}

// GetAllUsers
func (s *UserSuite) TestGetAllUsers1(t provider.T) {
	t.Title("[GetAllUsers] Repository error")
	t.Tags("user", "get_all_users")
	t.Parallel()
	t.WithNewStep("Incorrect: repository error", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		userMockRepo := mocks.NewIUserRepository(t)

		userMockRepo.
			On("GetAllUsers", ctx).
			Return(nil, fmt.Errorf("get all users fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx)

		users, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).GetAllUsers(ctx)

		sCtx.Assert().Nil(users)
		sCtx.Assert().Error(err)
	})
}

func (s *UserSuite) TestGetAllUsers2(t provider.T) {
	t.Title("[GetAllUsers] Correct")
	t.Tags("user", "get_all_users")
	t.Parallel()
	t.WithNewStep("Correct", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		userMockRepo := mocks.NewIUserRepository(t)

		users := utils.UserObjectMother{}.DefaultUsers()

		userMockRepo.
			On("GetAllUsers", ctx).
			Return(users, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx)

		users, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).GetAllUsers(ctx)

		sCtx.Assert().NotNil(users)
		sCtx.Assert().Nil(err)
	})
}

// UpdateUser
func (s *UserSuite) TestGetUpdateUser1(t provider.T) {
	t.Title("[UpdateUser] Incorrrect ID")
	t.Tags("user", "update_user")
	t.Parallel()
	t.WithNewStep("Incorrect: incorrect ID", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.IncorrectUserIDToUpdate()
		userMockRepo := mocks.NewIUserRepository(t)

		userMockRepo.
			On("GetUserByID", ctx, req.ID).
			Return(nil, fmt.Errorf("get user by id fail")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).UpdateRole(ctx, req)

		sCtx.Assert().Nil(user)
		sCtx.Assert().Error(err)
	})
}

func (s *UserSuite) TestGetUpdateUser2(t provider.T) {
	t.Title("[UpdateUser] Correct")
	t.Tags("user", "update_user")
	t.Parallel()
	t.WithNewStep("Correct: correct request", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.CorrectUserToUpdate()
		userMockRepo := mocks.NewIUserRepository(t)

		userCustomer := utils.UserObjectMother{}.DefaultCustomer()
		userAdmin := utils.UserObjectMother{}.DefaultAdmin()

		userMockRepo.
			On("GetUserByID", ctx, req.ID).
			Return(userCustomer, nil).
			Once()

		userMockRepo.
			On("UpdateRole", ctx, userAdmin).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).UpdateRole(ctx, req)

		sCtx.Assert().NotNil(user)
		sCtx.Assert().Nil(err)
	})
}
