package repository

import (
	"context"
	"gochi-boilerplate/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id, full_name, email, password, role, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.DB.Exec(ctx, query, user.ID, user.FullName, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var u model.User
	query := `SELECT id, full_name, email, password, role, created_at, updated_at 
			FROM users WHERE email = $1`
	err := r.DB.QueryRow(ctx, query, email).Scan(&u.ID, &u.FullName, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}