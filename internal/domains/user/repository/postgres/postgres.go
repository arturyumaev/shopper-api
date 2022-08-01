package postgres

import (
	"context"
	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/arturyumaev/shopper-api/models"
	"github.com/jackc/pgx/v4"
)

type repository struct {
	db *pgx.Conn
}

func (repo *repository) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (repo *repository) GetUser(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (repo *repository) GetUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}

func (repo *repository) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (repo *repository) DeleteUser(ctx context.Context, userId string) error {
	return nil
}

func NewRepository(db *pgx.Conn) user.Repository {
	return &repository{db}
}
