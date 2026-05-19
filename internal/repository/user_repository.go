package repository

import (
	"context"

	"github.com/ElCesarSP/go-api/internal/model"
	db "github.com/ElCesarSP/go-api/prisma/db"
)

type UserRepository struct {
	db *db.PrismaClient
}

func NewUserRepository(dbClient *db.PrismaClient) *UserRepository {
	return &UserRepository{db: dbClient}
}

func (r *UserRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.UserResponse, error) {
	user, err := r.db.User.CreateOne(
		db.User.Name.Set(input.Name),
		db.User.Email.Set(input.Email),
		db.User.Password.Set(input.Password),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*db.UserModel, error) {
	user, err := r.db.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id string,
) (*db.UserModel, error) {

	user, err := r.db.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
