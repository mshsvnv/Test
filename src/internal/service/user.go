package service

import (
	"context"
	"fmt"

	"src/internal/dto"
	"src/internal/model"
	repo "src/internal/repository"
	"src/pkg/logging"
	"src/pkg/utils"
)

//go:generate mockery --name=IUserService
type IUserService interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, req *dto.UpdateReq) (*model.User, error)
	UpdatePassword(ctx context.Context, req *dto.UpdatePasswordReq) (*model.User, error)
}

type UserService struct {
	logger logging.Interface
	repo   repo.IUserRepository
}

func NewUserService(
	logger logging.Interface,
	repo repo.IUserRepository) *UserService {
	return &UserService{
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) Update(ctx context.Context, req *dto.UpdateReq) (*model.User, error) {

	s.logger.Infof("update role id %s", req.ID)
	user, err := s.repo.GetUserByID(ctx, req.ID)

	if err != nil {
		s.logger.Errorf("get user by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by id fail, error %s", err)
	}

	user.Role = req.Role

	err = s.repo.Update(ctx, user)

	if err != nil {
		s.logger.Errorf("update fail, error %s", err.Error())
		return nil, fmt.Errorf("update fail, error %s", err)
	}

	return user, nil
}

func (s *UserService) UpdatePassword(ctx context.Context, req *dto.UpdatePasswordReq) (*model.User, error) {

	s.logger.Infof("update password id %s", req.ID)
	user, err := s.repo.GetUserByID(ctx, req.ID)

	if err != nil {
		s.logger.Errorf("get user by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by id fail, error %s", err)
	}

	user.Password = utils.HashAndSalt([]byte(req.Password))

	err = s.repo.Update(ctx, user)

	if err != nil {
		s.logger.Errorf("update fail, error %s", err.Error())
		return nil, fmt.Errorf("update fail, error %s", err)
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {

	s.logger.Infof("get all users")
	users, err := s.repo.GetAllUsers(ctx)

	if err != nil {
		s.logger.Errorf("get all users fail, error %s", err.Error())
		return nil, fmt.Errorf("get all users fail, error %s", err)
	}

	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {

	s.logger.Infof("get user by id %d", id)
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		s.logger.Errorf("get user by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by id fail, error %s", err)
	}

	return user, nil
}
