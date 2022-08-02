package postgres

import (
	"context"

	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/arturyumaev/shopper-api/internal/domains/user/queries"
	"github.com/arturyumaev/shopper-api/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type repository struct {
	db *pgx.Conn
}

func (repo *repository) CreateUser(ctx context.Context, user *models.User) error {
	user.ID = uuid.New().String()
	_, err := repo.db.Exec(ctx, queries.CreateUser, user.ID, user.Username, user.Email, user.Password)
	return err
}

func (repo *repository) GetUser(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (repo *repository) GetUsers(ctx context.Context) ([]*models.User, error) {
	rows, err := repo.db.Query(ctx, queries.SelectUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
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
