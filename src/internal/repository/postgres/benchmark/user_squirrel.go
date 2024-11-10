package benchmark

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"

	"src/internal/model"
)

type IUserRepositorySquirrel interface {
	Create(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type UserRepositorySquirrel struct {
	db *sql.DB
}

func NewUserRepositorySquirrel(db *sql.DB) IUserRepositorySquirrel {
	return &UserRepositorySquirrel{db}
}

func (r *UserRepositorySquirrel) Create(ctx context.Context, user *model.User) error {

	query := squirrel.
		Insert("user").
		Columns(
			"name",
			"surname",
			"email",
			"password",
			"role").
		Values(user.Name,
			user.Surname,
			user.Email,
			user.Password,
			user.Role).
		Suffix("returning id")

	sql, ars, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.db.QueryRow(sql, ars...)

	err = row.Scan(
		&user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositorySquirrel) GetUserByID(ctx context.Context, id int) (*model.User, error) {

	query := squirrel.
		Select("*").
		From("user").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sql, args...)

	var user model.User
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
