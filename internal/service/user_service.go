package service

import (
	"context"

	"github.com/ElCesarSP/go-api/internal/model"
	"github.com/ElCesarSP/go-api/internal/repository"
	"github.com/ElCesarSP/go-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	input model.CreateUserInput,
) (*model.UserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	input.Password = string(hashedPassword)

	return s.repo.Create(ctx, input)
}

func (s *UserService) Login(
	ctx context.Context,
	input model.LoginInput,
) (string, error) {

	user, err := s.repo.FindByEmail(ctx, input.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Me(
	ctx context.Context,
	userID string,
) (*model.UserResponse, error) {

	user, err := s.repo.FindByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
