package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(float64)

	var user models.Users
	if err := configs.DB.First(&user, uint(userID)).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	helpers.SuccessResponse(c, "User profile fetched successfully", gin.H{
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
		"created_at":        user.CreatedAt,
		"updated_at":        user.UpdatedAt,
	})
}

func GetAllUsers(c *gin.Context){
	var users []models.Users
				if err := configs.DB.Find(&users).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve users")
        return
    }

    helpers.SuccessResponse(c, "Users retrieved successfully", users)
}
