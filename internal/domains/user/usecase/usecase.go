package usecase

import (
	"context"
	"errors"

	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/arturyumaev/shopper-api/models"
)

type useCase struct {
	repository user.Repository
}

func (uc *useCase) CreateUser(ctx context.Context, user *models.User) error {
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	err := uc.repository.CreateUser(ctx, user)
	return err
}

func (uc *useCase) GetUser(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (uc *useCase) GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := uc.repository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *useCase) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (uc *useCase) DeleteUser(ctx context.Context, userId string) error {
	return nil
}

func NewUseCase(repository user.Repository) user.UseCase {
	return &useCase{repository}
}
