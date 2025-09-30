package repository

import (
	"context"
	"gochi-boilerplate/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	DB *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO products (id, name, price, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.Exec(ctx, query, product.ID, product.Name, product.Price, product.UserID, product.CreatedAt, product.UpdatedAt)
	return err
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{}
	query := `SELECT id, name, price, created_at, updated_at FROM products`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	var p model.Product
	query := `SELECT id, name, price, created_at, updated_at FROM products WHERE id = $1`
	err := r.DB.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product *model.Product) error {
	query := `UPDATE products SET name = $1, price = $2, updated_at = $3 WHERE id = $4`
	_, err := r.DB.Exec(ctx, query, product.Name, product.Price, product.UpdatedAt, product.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}