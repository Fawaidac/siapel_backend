package backoffice

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/Fawaidac/siapel_backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func CreateKKBackoffice(c *gin.Context) {
	// Retrieve user_id from context
	userID, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "User ID not found in token")
		return
	}

	userIDUint, ok := userID.(float64)
	if !ok {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid user ID type")
		return
	}

	// Bind JSON request
	var request struct {
		ServiceID  uint   `json:"service_id" binding:"required"`
		Name       string `json:"name" binding:"required"`
		NIK        string `json:"nik" binding:"required"`
		NoTelp     string `json:"no_telp" binding:"required"`
		KK         string `json:"kk"`
		Akta       string `json:"akta"`
		FotoKTP    string `json:"foto_ktp"`
		Kecamatan  string `json:"kecamatan" binding:"required"`
		Kelurahan  string `json:"kelurahan" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Fetch service details
	var service models.Service
	if err := configs.DB.Where("id = ?", request.ServiceID).First(&service).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Service not found")
		} else {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Database error")
		}
		return
	}

	// Generate operator ID
	operatorID, err := utils.GenerateOperator(configs.DB, "OP_KK")
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate operator")
		return
	}

	// Generate ticket code
	ticketCode := utils.GenerateTicketString(service.Code)

	// Start transaction
	tx := configs.DB.Begin()
	if tx.Error != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to begin transaction")
		return
	}

	// Create registration entry
	registration := models.Registration{
		UserID:     uint(userIDUint),
		OperatorID: operatorID,
		Tiket:      ticketCode,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := tx.Create(&registration).Error; err != nil {
		tx.Rollback()
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create registration")
		return
	}

	// Insert requirements based on service ID
	var requirements []models.Requirement
	if request.ServiceID == 1 {
		requirements = []models.Requirement{
			{Name: "Name", Data: request.Name, IsValid: "0"},
			{Name: "NIK", Data: request.NIK, IsValid: "0"},
			{Name: "NoTelp", Data: request.NoTelp, IsValid: "0"},
			{Name: "Kecamatan", Data: request.Kecamatan, IsValid: "0"},
			{Name: "Kelurahan", Data: request.Kelurahan, IsValid: "0"},
			{Name: "KK", Data: request.KK, IsValid: "0"},
			{Name: "Akta Kelahiran", Data: request.Akta, IsValid: "0"},
		}
	} else if request.ServiceID == 2 {
		requirements = []models.Requirement{
			{Name: "Name", Data: request.Name, IsValid: "0"},
			{Name: "NIK", Data: request.NIK, IsValid: "0"},
			{Name: "NoTelp", Data: request.NoTelp, IsValid: "0"},
			{Name: "Kecamatan", Data: request.Kecamatan, IsValid: "0"},
			{Name: "Kelurahan", Data: request.Kelurahan, IsValid: "0"},
			{Name: "KK", Data: request.KK, IsValid: "0"},
			{Name: "Foto KTP", Data: request.FotoKTP, IsValid: "0"},
		}
	} else {
		tx.Rollback()
		helpers.ErrorResponse(c, http.StatusBadRequest, "Unsupported service ID")
		return
	}

	for i, req := range requirements {
		if req.Name == "KK" || req.Name == "Akta Kelahiran" || req.Name == "Foto KTP" {
			uploadPath, err := utils.HandleFileUpload(c, req.Name, filepath.Join("persyaratan", req.Name))
			if err != nil {
				tx.Rollback()
				helpers.ErrorResponse(c, http.StatusInternalServerError, "File upload error")
				return
			}
			requirements[i].Data = uploadPath
		}

		if err := tx.Create(&requirements[i]).Error; err != nil {
			tx.Rollback()
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to insert requirement")
			return
		}
	}

	// Insert registration details
	for _, req := range requirements {
		detail := models.RegistrationDetail{
			IDPendaftaran: registration.ID,
			IDPersyaratan: req.ID,
			IDLayanan:     request.ServiceID,
		}
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to insert registration detail")
			return
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to commit transaction")
		return
	}

	// Send success response
	helpers.SuccessResponse(c, "Registration created successfully", gin.H{
		"ticket":       registration.Tiket,
		"registration": registration,
	})
}
