package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims adalah struktur custom untuk data di dalam token JWT
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken membuat token JWT baru untuk pengguna
func GenerateToken(userID, role string) (string, error) {
	jwtSecret := GetEnv("JWT_SECRET", "supersecret")
	
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token berlaku 24 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken memvalidasi token JWT dan mengembalikan claims
func ValidateToken(tokenString string) (*Claims, error) {
	jwtSecret := GetEnv("JWT_SECRET", "supersecret")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, fmt.Errorf("invalid token")
}