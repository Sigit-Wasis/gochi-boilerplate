package middleware

import (
	"context"
	"gochi-boilerplate/internal/utils"
	"net/http"
	"strings"
)

// ContextKey adalah tipe custom untuk kunci context agar tidak bentrok
type ContextKey string

const UserClaimsKey ContextKey = "userClaims"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, "Header Authorization dibutuhkan", "missing auth header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.RespondError(w, http.StatusUnauthorized, "Format header Authorization salah", "format must be Bearer <token>")
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Token tidak valid", err.Error())
			return
		}

		// Simpan claims di context agar bisa diakses oleh handler selanjutnya
		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}