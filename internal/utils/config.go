package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig memuat variabel lingkungan dari file .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: file .env tidak ditemukan, menggunakan variabel lingkungan sistem")
	}
}

// GetEnv mengambil variabel lingkungan atau mengembalikan nilai default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}