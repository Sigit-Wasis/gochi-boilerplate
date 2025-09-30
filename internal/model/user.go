package model

import (
	"time"

	"github.com/google/uuid"
)

// User struct sesuai dengan tabel 'users' di database
type User struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Jangan pernah kirim password hash ke klien
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RegisterRequest adalah model untuk body request registrasi
type RegisterRequest struct {
	FullName string `json:"full_name" example:"John Doe"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" example:"OnlinePHP"`
}

// LoginRequest adalah model untuk body request login
type LoginRequest struct {
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" example:"OnlinePHP"`
}

// LoginResponse adalah model untuk respon setelah login sukses
type LoginResponse struct {
	Token string `json:"token"`
}