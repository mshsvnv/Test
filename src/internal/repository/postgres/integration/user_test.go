//go:build integration

package mypostgres_test

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	repo "src/internal/repository"
	"src/internal/repository/utils"
)

type UserRepoSuite struct {
	suite.Suite

	userRepo repo.IUserRepository
	userID   int
}

// func (u *UserRepoSuite) TestIntegrationUserRepoCreate(t provider.T) {
// 	t.Title("[Create] Create user")
// 	t.Tags("integration", "user", "repository", "postgres")
// 	t.Parallel()
// 	t.WithNewStep("Create user", func(sCtx provider.StepCtx) {

// 		ctx := context.TODO()
// 		request := utils.UserObjectMother{}.
// 			WithName("Dmitriy").
// 			WithSurname("Dmitrov").
// 			WithPassword("dmitry").
// 			WithEmail("dmitry@mail.ru").
// 			WithRole(model.UserRoleCustomer).
// 			ToModel()

// 		sCtx.WithNewParameters("ctx", ctx, "request", request)

// 		err := u.userRepo.Create(ctx, request)
// 		sCtx.Assert().NoError(err)

// 		err = u.userRepo.Delete(ctx, request.ID)
// 		sCtx.Assert().NoError(err)
// 	})
// }

func (u *UserRepoSuite) TestIntegrationUserRepoUpdateRole(t provider.T) {
	t.Title("[Update] Update user role")
	t.Tags("integration", "user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Update user role", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request, _ := u.userRepo.GetUserByID(ctx, u.userID)

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := u.userRepo.UpdateRole(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (u *UserRepoSuite) TestIntegrationUserRepoGetAllUsers(t provider.T) {
	t.Title("[GetAllUsers] Get all users")
	t.Tags("integration", "user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get all users", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user, _ := u.userRepo.GetUserByID(ctx, u.userID)

		sCtx.WithNewParameters("ctx", ctx)

		users, err := u.userRepo.GetAllUsers(ctx)

		sCtx.Assert().NotEmpty(users)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(users[0], user)
	})
}

func (u *UserRepoSuite) TestIntegrationUserRepoGetUserByID(t provider.T) {
	t.Title("[GetUserByID] Get user by id")
	t.Tags("integration", "user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get user by id", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "request", u.userID)

		user, err := u.userRepo.GetUserByID(ctx, u.userID)

		sCtx.Assert().NotEmpty(user)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user.ID, u.userID)
		sCtx.Assert().Equal(user.Name, "Ivan")
		sCtx.Assert().Equal(user.Surname, "Ivanov")
		sCtx.Assert().Equal(user.Email, "ivan@mail.ru")
	})
}

func (u *UserRepoSuite) TestIntegrationUserRepoGetUserByEmail(t provider.T) {
	t.Title("[GetUserByEmail] Get user by email")
	t.Tags("integration", "user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Get user by email", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		req := utils.UserObjectMother{}.CorrectEmail()
		userTmp, _ := u.userRepo.GetUserByID(ctx, u.userID)

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		user, err := u.userRepo.GetUserByEmail(ctx, req)

		sCtx.Assert().NotEmpty(user)
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, userTmp)
	})
}

func (u *UserRepoSuite) TestIntegrationUserRepoDelete(t provider.T) {
	t.Title("[Delete] Delete user")
	t.Tags("integration", "user", "repository", "postgres")
	t.Parallel()
	t.WithNewStep("Delete user", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.UserObjectMother{}.
			WithName("Misha").
			WithSurname("Mihailov").
			ToModel()

		u.userRepo.Create(ctx, user)
		req := user.ID

		sCtx.WithNewParameters("ctx", ctx, "request", req)

		err := u.userRepo.Delete(ctx, req)

		sCtx.Assert().NoError(err)
	})
}
