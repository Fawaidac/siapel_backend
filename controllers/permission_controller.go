package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePermission(c *gin.Context) {
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := configs.DB.Create(&permission).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to create permission")
		return
	}

	helpers.SuccessResponse(c, "Permission created successfully", permission)
}

func GetPermissions(c *gin.Context) {
	var permissions []models.Permission
	if err := configs.DB.Find(&permissions).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch permissions")
		return
	}

	helpers.SuccessResponse(c, "Permissions fetched successfully", permissions)
}

func GetPermission(c *gin.Context) {
	permissionID := c.Param("id")
	var permission models.Permission
	if err := configs.DB.First(&permission, permissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Permission not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch permission")
		return
	}

	helpers.SuccessResponse(c, "Permission fetched successfully", permission)
}

func UpdatePermission(c *gin.Context) {
	permissionID := c.Param("id")
	var permission models.Permission
	if err := configs.DB.First(&permission, permissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Permission not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch permission")
		return
	}

	if err := c.ShouldBindJSON(&permission); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := configs.DB.Save(&permission).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to update permission")
		return
	}

	helpers.SuccessResponse(c, "Permission updated successfully", permission)
}

func DeletePermission(c *gin.Context) {
	permissionID := c.Param("id")
	var permission models.Permission
	if err := configs.DB.First(&permission, permissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Permission not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch permission")
		return
	}

	if err := configs.DB.Delete(&permission).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to delete permission")
		return
	}

	helpers.SuccessResponse(c, "Permission deleted successfully", nil)
}