package main

import (
	"context"
	"fmt"
	"gochi-boilerplate/internal/handler"
	"gochi-boilerplate/internal/middleware"
	"gochi-boilerplate/internal/repository"
	"gochi-boilerplate/internal/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "gochi-boilerplate/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Boilerplate API with Go, Chi, PostgreSQL, and Swagger
// @version 1.0
// @description This is a sample server for a Go CRUD application with Chi, PostgreSQL, and Swagger.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email hellowasis@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	utils.LoadConfig()
	dbURL := utils.GetEnv("DATABASE_URL", "")
	port := utils.GetEnv("SERVER_PORT", "8080")

	if dbURL == "" {
		log.Fatal("DATABASE_URL harus diatur di environment variable")
	}

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Tidak bisa terhubung ke database: %v\n", err)
	}
	defer dbpool.Close()

	// 2. Inisialisasi Repository dan Handler baru untuk User & Auth
	userRepo := repository.NewUserRepository(dbpool)
	productRepo := repository.NewProductRepository(dbpool)

	authHandler := handler.NewAuthHandler(userRepo)
	productHandler := handler.NewProductHandler(productRepo)

	r := chi.NewRouter()

	// Middleware Global
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Rute Swagger (Publik)
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))

	// 3. Rute Publik untuk Autentikasi
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	// 4. Grup Rute Terproteksi yang memerlukan JWT
	r.Group(func(r chi.Router) {
		// Gunakan AuthMiddleware di sini untuk melindungi semua rute di dalam grup ini
		r.Use(middleware.AuthMiddleware)

		// Rute untuk produk sekarang berada di dalam grup yang dilindungi
		r.Route("/products", func(r chi.Router) {
			r.Post("/", productHandler.CreateProduct)
			r.Get("/", productHandler.GetAllProducts)
			r.Get("/{id}", productHandler.GetProductByID)
			r.Put("/{id}", productHandler.UpdateProduct)
			r.Delete("/{id}", productHandler.DeleteProduct)
		})
	})

	// Menjalankan Server
	fmt.Printf("Server berjalan di port %s\n", port)
	fmt.Printf("Swagger UI tersedia di http://localhost:%s/swagger/index.html\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}