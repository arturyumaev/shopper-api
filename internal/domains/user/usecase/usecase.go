package usecase

import (
	"context"
	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/arturyumaev/shopper-api/models"
)

type useCase struct {
	repository user.Repository
}

func (uc *useCase) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (uc *useCase) GetUser(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (uc *useCase) GetUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
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
