package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"go_health_sentiment/auth"
	"go_health_sentiment/models"
	"go_health_sentiment/utils"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string                `json:"token"`
	User  models.UserResponse   `json:"user"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate input
	var validationErrors []utils.ValidationError
	
	req.Email = utils.SanitizeInput(req.Email)
	if !utils.ValidateEmail(req.Email) {
		validationErrors = append(validationErrors, utils.ValidationError{
			Field:   "email",
			Message: "Invalid email format",
		})
	}

	passwordErrors := utils.ValidatePassword(req.Password)
	validationErrors = append(validationErrors, passwordErrors...)

	if len(validationErrors) > 0 {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error processing password")
		return
	}

	// Create user
	user := models.User{
		Email:    strings.ToLower(req.Email),
		Password: string(hashedPassword),
	}

	if err := user.CreateUser(h.db); err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			utils.WriteError(w, http.StatusConflict, "Email already registered")
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	response := AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}

	utils.WriteCreated(w, "User registered successfully", response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate input
	req.Email = utils.SanitizeInput(req.Email)
	if !utils.ValidateEmail(req.Email) {
		utils.WriteError(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	// Get user by email
	user, err := models.GetUserByEmail(h.db, strings.ToLower(req.Email))
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	response := AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}

	utils.WriteSuccess(w, "Login successful", response)
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserKey).(int)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	user, err := models.GetUserByID(h.db, userID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.WriteSuccess(w, "Profile retrieved successfully", user.ToResponse())
}