package benchmark

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"src/internal/model"
)

type IUserRepositoryGorm interface {
	Create(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) IUserRepositoryGorm {
	return &UserRepositoryGorm{db}
}

func (r *UserRepositoryGorm) Create(ctx context.Context, user *model.User) error {

	res := r.db.Table("user").Create(user)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}

	return nil
}

func (r *UserRepositoryGorm) GetUserByID(ctx context.Context, id int) (*model.User, error) {

	user := &model.User{}

	res := r.db.Table("user").Where("id = ?", id).Take(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
