package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name  string `json:"name" example:"Laptop Gaming"`
	Price int    `json:"price" example:"15000000"`
}

type UpdateProductRequest struct {
	Name  *string `json:"name,omitempty"`
	Price *int    `json:"price,omitempty"`
}