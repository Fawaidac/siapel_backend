package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		IDKecamatan uint   `json:"id_kecamatan" binding:"required"`
		IDKelurahan uint   `json:"id_kelurahan" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=6"`
		NIK         string `json:"nik" binding:"required"`
		Phone       string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.Users{
		IDKecamatan: input.IDKecamatan,
		IDKelurahan: input.IDKelurahan,
		Name:        input.Name,
		Email:       input.Email,
		Password:    string(hashedPassword),
		NIK:         input.NIK,
		Phone:       input.Phone,
		Status:      "active",
	}

	if err := configs.DB.Where("email = ?", input.Email).Or("nik = ?", input.NIK).First(&models.Users{}).Error; err == nil {
	helpers.ErrorResponse(c, http.StatusConflict, "Email or NIK already registered")
	return
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	helpers.SuccessResponse(c, "User registered successfully", gin.H{"id": user.ID})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.Users
	if err := configs.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := configs.GenerateJWT(user.ID, user.Email)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	helpers.SuccessResponse(c, "Login successful", gin.H{
		"user": gin.H{
			"id":                user.ID,
			"id_kecamatan":      user.IDKecamatan,
			"id_kelurahan":      user.IDKelurahan,
			"name":              user.Name,
			"email":             user.Email,
			"email_verified_at": user.EmailVerifiedAt,
			"nik":               user.NIK,
			"phone":             user.Phone,
			"avatar":            user.Avatar,
			"hit":               user.Hit,
			"status":            user.Status,
		},
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	
	helpers.SuccessResponse(c, "Logged out successfully", gin.H{})
}