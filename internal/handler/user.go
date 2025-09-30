package handler

import (
	"encoding/json"
	"gochi-boilerplate/internal/model"
	"gochi-boilerplate/internal/repository"
	"gochi-boilerplate/internal/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: userRepo}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with full name, email, and password.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        user body model.RegisterRequest true "User Registration Details"
// @Success      201  {object}  utils.Response "Successfully registered"
// @Failure      400  {object}  utils.Response "Invalid request body"
// @Failure      500  {object}  utils.Response "Internal server error"
// @Router       /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Request body tidak valid", err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal memproses password", err.Error())
		return
	}

	user := &model.User{
		ID:        uuid.New(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.UserRepo.CreateUser(r.Context(), user); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal membuat pengguna", err.Error())
		return
	}

	utils.RespondSuccess(w, http.StatusCreated, "Registrasi berhasil", nil)
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticate a user with email and password to get a JWT.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        credentials body model.LoginRequest true "User Login Credentials"
// @Success      200  {object}  utils.Response{data=model.LoginResponse} "Successfully logged in with token"
// @Failure      400  {object}  utils.Response "Invalid request body"
// @Failure      401  {object}  utils.Response "Unauthorized - Invalid credentials"
// @Failure      500  {object}  utils.Response "Internal server error"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Request body tidak valid", err.Error())
		return
	}

	user, err := h.UserRepo.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "Email atau password salah", "user not found")
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.RespondError(w, http.StatusUnauthorized, "Email atau password salah", "invalid password")
		return
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Gagal membuat token", err.Error())
		return
	}
	
	resp := model.LoginResponse{Token: token}
	utils.RespondSuccess(w, http.StatusOK, "Login berhasil", resp)
}