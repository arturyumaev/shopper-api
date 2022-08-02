package user

import (
	"context"

	"github.com/arturyumaev/shopper-api/models"
)

type UseCase interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, userId string) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, userId string, user *models.User) error
	DeleteUser(ctx context.Context, userId string) error
}
