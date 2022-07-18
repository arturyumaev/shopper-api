package user

import (
	"context"

	"github.com/arturyumaev/shopper-api/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, userId string) error
}
