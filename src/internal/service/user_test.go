//go:build unit

package service_test

import (
	"context"
	"fmt"
	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/service"
	"src/internal/service/utils"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UserServiceSuite struct {
	suite.Suite
}

// GetUserByID
func (s *UserServiceSuite) TestUserServiceGetByID1(t provider.T) {
	t.Title("[GetUserByID] Incorrect ID")
	t.Tags("user", "service", "service", "get_user_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: not existed user", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.IncorrectID()
		userMockRepo := mocks.NewIUserRepository(t)

		userMockRepo.
			On("GetUserByID", ctx, req).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).GetUserByID(ctx, req)

		sCtx.Assert().Nil(user)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *UserServiceSuite) TestUserServiceGetUserByID2(t provider.T) {
	t.Title("[GetUserByID] Correct ID")
	t.Tags("user", "service", "service", "get_user_by_id")
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
func (s *UserServiceSuite) TestUserServiceGetAllUsers1(t provider.T) {
	t.Title("[GetAllUsers] Repository error")
	t.Tags("user", "service", "service", "get_all_users")
	t.Parallel()
	t.WithNewStep("Incorrect: repository error", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		userMockRepo := mocks.NewIUserRepository(t)

		userMockRepo.
			On("GetAllUsers", ctx).
			Return(nil, fmt.Errorf("no rows in result set")).
			Once()

		sCtx.WithNewParameters("ctx", ctx)

		users, err := service.NewUserService(utils.NewMockLogger(), userMockRepo).GetAllUsers(ctx)

		sCtx.Assert().Nil(users)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *UserServiceSuite) TestUserServiceGetAllUsers2(t provider.T) {
	t.Title("[GetAllUsers] Correct")
	t.Tags("user", "service", "service", "get_all_users")
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
