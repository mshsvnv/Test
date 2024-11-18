//go:build unit

package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
	"src/internal/repository/utils"
)

type UserRepoSuite struct {
	suite.Suite

	userMockRepo mocks.IUserRepository
}

func (u *UserRepoSuite) BeforeAll(t provider.T) {
	t.Title("Init user mock repo")
	u.userMockRepo = *mocks.NewIUserRepository(t)
	t.Tags("fixture", "user")
}

func (u *UserRepoSuite) TestUserRepoCreate(t provider.T) {
	t.Title("[Create] Create user")
	t.Tags("user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Create user", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.UserObjectMother{}.DefaultCustomer(1)

		u.userMockRepo.
			On("Create", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := u.userMockRepo.Create(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (u *UserRepoSuite) TestUserRepoUpdate(t provider.T) {
	t.Title("[Update] Update user role")
	t.Tags("user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update user role", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.UserObjectMother{}.DefaultCustomer(1)

		u.userMockRepo.
			On("Update", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := u.userMockRepo.Update(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (u *UserRepoSuite) TestUserRepoGetAllUsers(t provider.T) {
	t.Title("[GetAllUsers] Get all users")
	t.Tags("user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get all users", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		expUsers := []*model.User{
			{
				ID:       1,
				Name:     "Stepan",
				Surname:  "Postnov",
				Email:    "pstpn@gmail.com",
				Password: "1",
				Role:     model.UserRoleAdmin,
			},
		}

		u.userMockRepo.
			On("GetAllUsers", ctx).
			Return(expUsers, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx)

		users, err := u.userMockRepo.GetAllUsers(ctx)

		sCtx.Assert().NotEmpty(users)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(users, expUsers)
	})
}

func (u *UserRepoSuite) TestUserRepoGetUserByID(t provider.T) {
	t.Title("[GetUserByID] Get user by id")
	t.Tags("user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get user by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		req := utils.UserObjectMother{}.CorrectID()
		expUser := &model.User{
			ID:       req,
			Name:     "Stepan",
			Surname:  "Postnov",
			Email:    "pstpn@gmail.com",
			Password: "1",
			Role:     model.UserRoleAdmin,
		}

		u.userMockRepo.
			On("GetUserByID", ctx, req).
			Return(expUser, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := u.userMockRepo.GetUserByID(ctx, req)

		sCtx.Assert().NotEmpty(user)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, expUser)
	})
}

func (u *UserRepoSuite) TestUserRepoGetUserByEmail(t provider.T) {
	t.Title("[GetUserByEmail] Get user by email")
	t.Tags("user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get user by email", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		req := utils.UserObjectMother{}.CorrectEmail()
		expUser := &model.User{
			ID:       1,
			Name:     "Stepan",
			Surname:  "Postnov",
			Email:    req,
			Password: "1",
			Role:     model.UserRoleAdmin,
		}

		u.userMockRepo.
			On("GetUserByEmail", ctx, req).
			Return(expUser, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := u.userMockRepo.GetUserByEmail(ctx, req)

		sCtx.Assert().NotEmpty(user)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, expUser)
	})
}

func (u *UserRepoSuite) TestUserRepoDelete(t provider.T) {
	t.Title("[Delete] Delete user")
	t.Tags("user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Delete user", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		req := utils.UserObjectMother{}.CorrectID()
		u.userMockRepo.
			On("Delete", ctx, req).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := u.userMockRepo.Delete(ctx, req)

		sCtx.Assert().NoError(err)
	})
}
