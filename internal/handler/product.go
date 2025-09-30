package handler

import (
	"encoding/json"
	"gochi-boilerplate/internal/middleware"
	"gochi-boilerplate/internal/model"
	"gochi-boilerplate/internal/repository"
	"gochi-boilerplate/internal/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Add a new product to the database. The product will be associated with the logged-in user.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        product body model.CreateProductRequest true "Create Product"
// @Success      201  {object}  utils.Response{data=model.Product}
// @Failure      400  {object}  utils.Response "Bad Request"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Ambil claims pengguna dari context yang sudah diisi oleh middleware
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*utils.Claims)
	if !ok {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mendapatkan data pengguna dari token", "invalid context claims")
		return
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal memproses ID pengguna", err.Error())
		return
	}

	var req model.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Request body tidak valid", err.Error())
		return
	}

	product := &model.Product{
		ID:        uuid.New(),
		Name:      req.Name,
		Price:     req.Price,
		UserID:    &userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.Repo.CreateProduct(r.Context(), product); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal membuat produk", err.Error())
		return
	}

	utils.RespondSuccess(w, http.StatusCreated, "Produk berhasil dibuat", product)
}

// GetAllProducts godoc
// @Summary      Get all products
// @Description  Get a list of all products. Requires authentication.
// @Tags         Products
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  utils.Response{data=[]model.Product}
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repo.GetAllProducts(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mengambil semua produk", err.Error())
		return
	}
	utils.RespondSuccess(w, http.StatusOK, "Berhasil mengambil semua produk", products)
}

// GetProductByID godoc
// @Summary      Get a product by ID
// @Description  Get a single product by its UUID. Requires authentication.
// @Tags         Products
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Product ID" format(uuid)
// @Success      200  {object}  utils.Response{data=model.Product}
// @Failure      400  {object}  utils.Response "Invalid UUID format"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      404  {object}  utils.Response "Product not found"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Format UUID tidak valid", err.Error())
		return
	}

	product, err := h.Repo.GetProductByID(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Produk tidak ditemukan", err.Error())
		return
	}
	utils.RespondSuccess(w, http.StatusOK, "Berhasil menemukan produk", product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update an existing product's details. Only the product owner or an admin can perform this action.
// @Tags         Products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Product ID" format(uuid)
// @Param        product body model.UpdateProductRequest true "Update Product"
// @Success      200  {object}  utils.Response{data=model.Product}
// @Failure      400  {object}  utils.Response "Bad Request"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      403  {object}  utils.Response "Forbidden"
// @Failure      404  {object}  utils.Response "Product not found"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Ambil claims pengguna dari context
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*utils.Claims)
	if !ok {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mendapatkan data pengguna dari token", "invalid context claims")
		return
	}
	userIDFromToken, _ := uuid.Parse(claims.UserID)

	// Ambil ID produk dari URL
	productIDStr := chi.URLParam(r, "id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Format UUID produk tidak valid", err.Error())
		return
	}

	// Ambil produk yang ada dari database
	existingProduct, err := h.Repo.GetProductByID(r.Context(), productID)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Produk tidak ditemukan", err.Error())
		return
	}

	// Otorisasi: Cek apakah pengguna adalah pemilik produk atau seorang admin
	if (existingProduct.UserID == nil || *existingProduct.UserID != userIDFromToken) && claims.Role != "admin" {
		utils.RespondError(w, http.StatusForbidden, "Akses ditolak", "Anda tidak memiliki izin untuk mengubah produk ini")
		return
	}

	var req model.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Request body tidak valid", err.Error())
		return
	}
	if req.Name != nil {
		existingProduct.Name = *req.Name
	}
	if req.Price != nil {
		existingProduct.Price = *req.Price
	}
	existingProduct.UpdatedAt = time.Now()
	if err := h.Repo.UpdateProduct(r.Context(), existingProduct); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mengupdate produk", err.Error())
		return
	}

	utils.RespondSuccess(w, http.StatusOK, "Produk berhasil diupdate", existingProduct)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by its UUID. Only the product owner or an admin can perform this action.
// @Tags         Products
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Product ID" format(uuid)
// @Success      200  {object}  utils.Response "Successfully deleted"
// @Failure      400  {object}  utils.Response "Invalid UUID format"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      403  {object}  utils.Response "Forbidden"
// @Failure      404  {object}  utils.Response "Product not found"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Mirip dengan Update, kita cek kepemilikan
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*utils.Claims)
	if !ok {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mendapatkan data pengguna dari token", "invalid context claims")
		return
	}
	userIDFromToken, _ := uuid.Parse(claims.UserID)

	productIDStr := chi.URLParam(r, "id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Format UUID produk tidak valid", err.Error())
		return
	}

	// Cek kepemilikan sebelum menghapus
	product, err := h.Repo.GetProductByID(r.Context(), productID)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Produk tidak ditemukan", err.Error())
		return
	}

	if product.UserID == nil || *product.UserID != userIDFromToken && claims.Role != "admin" {
		utils.RespondError(w, http.StatusForbidden, "Akses ditolak", "Anda tidak memiliki izin untuk menghapus produk ini")
		return
	}

	if err := h.Repo.DeleteProduct(r.Context(), productID); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal menghapus produk", err.Error())
		return
	}

	utils.RespondSuccess(w, http.StatusOK, "Produk berhasil dihapus", nil)
}