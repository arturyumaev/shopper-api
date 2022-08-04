package usecase

import (
	"context"

	"github.com/arturyumaev/shopper-api/internal/domains/auth"
	"github.com/arturyumaev/shopper-api/models"
)

type useCase struct {
}

func (uc *useCase) SignUp(ctx context.Context, user *models.User) error {
	return nil
}

func (uc *useCase) SignIn(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}

func NewUseCase() auth.UseCase {
	return &useCase{}
}
