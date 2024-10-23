//go:build integration

package service_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/service"
	"src/internal/service/utils"
)

type UserSuite struct {
	suite.Suite

	userService service.IUserService
}

// GetUserByID
func (s *UserSuite) TestUserServiceGetByID1(t provider.T) {
	t.Title("[Integration GetUserByID] Incorrect ID")
	t.Tags("integration", "user", "service", "service", "get_user_by_id")
	t.Parallel()
	t.WithNewStep("Incorrect: not existed user", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.IncorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := s.userService.GetUserByID(ctx, req)

		sCtx.Assert().Nil(user)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *UserSuite) TestUserServiceGetUserByID2(t provider.T) {
	t.Title("[Integration GetUserByID] Correct ID")
	t.Tags("integration", "user", "service", "service", "get_user_by_id")
	t.Parallel()
	t.WithNewStep("Correct: existed user", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.CorrectID()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := s.userService.GetUserByID(ctx, req)

		sCtx.Assert().NotEmpty(user)
		sCtx.Assert().Nil(err)
	})
}

// GetAllUsers
func (s *UserSuite) TestUserServiceGetAllUsers2(t provider.T) {
	t.Title("[Integration GetAllUsers] Correct")
	t.Tags("integration", "user", "service", "service", "get_all_users")
	t.Parallel()
	t.WithNewStep("Correct", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		sCtx.WithNewParameters("ctx", ctx)

		users, err := s.userService.GetAllUsers(ctx)

		sCtx.Assert().NotNil(users)
		sCtx.Assert().Nil(err)
	})
}

// UpdateUser
func (s *UserSuite) TestUserServiceGetUpdateUser1(t provider.T) {
	t.Title("[Integration UpdateUser] Incorrrect ID")
	t.Tags("integration", "user", "service", "update_user")
	t.Parallel()
	t.WithNewStep("Incorrect: incorrect ID", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.IncorrectUserIDToUpdate()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := s.userService.UpdateRole(ctx, req)

		sCtx.Assert().Nil(user)
		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), pgx.ErrNoRows.Error())
	})
}

func (s *UserSuite) TestUserServiceGetUpdateUser2(t provider.T) {
	t.Title("[Integration UpdateUser] Correct")
	t.Tags("integration", "user", "service", "update_user")
	t.Parallel()
	t.WithNewStep("Correct: correct request", func(sCtx provider.StepCtx) {

		ctx := context.TODO()
		req := utils.UserObjectMother{}.CorrectUserToUpdate()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := s.userService.UpdateRole(ctx, req)

		sCtx.Assert().NotNil(user)
		sCtx.Assert().Nil(err)
	})
}
