package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword menghasilkan hash dari sebuah password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // Angka 14 adalah cost factor
	return string(bytes), err
}

// CheckPasswordHash membandingkan password dengan hash-nya
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // Jika tidak ada error, password cocok
}