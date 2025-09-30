# Makefile

# Nama aplikasi dan direktori output
APP_NAME=gochi-boilerplate
CMD_PATH=cmd/server/main.go
BIN_DIR=bin

# Variabel dari .env untuk koneksi database lokal
# Pastikan Anda sudah 'export' variabel ini atau gunakan 'source .env'
# Makefile tidak secara otomatis membaca .env, namun psql di bawah ini akan
# membaca DATABASE_URL jika aplikasi Go kita sudah mengekspornya.
# Cara paling mudah adalah memastikan DATABASE_URL di .env sudah benar.
include .env
export

# ====================================================================================
# APLIKASI GO 🚀
# ====================================================================================

.PHONY: run
run: ## Menjalankan aplikasi dalam mode pengembangan
	@echo "🔥 Running application..."
	@go run $(CMD_PATH)

.PHONY: build
build: ## Meng-kompilasi aplikasi menjadi binary
	@echo "📦 Building application binary..."
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)
	@echo "✅ Build complete: $(BIN_DIR)/$(APP_NAME)"

.PHONY: test
test: ## Menjalankan semua unit test
	@echo "🧪 Running tests..."
	@go test ./... -v

.PHONY: tidy
tidy: ## Merapikan dependensi go.mod
	@echo "🧹 Tidying go modules..."
	@go mod tidy

.PHONY: clean
clean: ## Menghapus hasil build
	@echo "🗑️ Cleaning up build artifacts..."
	@rm -rf $(BIN_DIR)

# ====================================================================================
# DOKUMENTASI SWAGGER 📄
# ====================================================================================

.PHONY: swag
swag: ## Membuat atau memperbarui dokumentasi Swagger
	@echo "📄 Generating Swagger docs..."
	@swag init -g $(CMD_PATH)
	@echo "✅ Swagger docs generated."

# ====================================================================================
# DATABASE (LOCAL POSTGRESQL) 🐘
# ====================================================================================

.PHONY: db-migrate
db-migrate: ## Menjalankan skrip SQL (migrasi) ke database lokal
	@echo "🐘 Running database migrations on local PostgreSQL..."
	@psql "$(DATABASE_URL)" -f ./db/migrations/001_init_schema.sql

.PHONY: db-connect
db-connect: ## Membuka terminal psql ke database lokal
	@echo "🐘 Connecting to local PostgreSQL shell..."
	@psql "$(DATABASE_URL)"


# ====================================================================================
# BANTUAN ℹ️
# ====================================================================================

.PHONY: help
help: ## Menampilkan daftar perintah ini
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'