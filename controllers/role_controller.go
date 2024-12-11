package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := configs.DB.Create(&role).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to create role")
		return
	}

	helpers.SuccessResponse(c, "Role created successfully", role)
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := configs.DB.Find(&roles).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch roles")
		return
	}

	helpers.SuccessResponse(c, "Roles fetched successfully", roles)
}

func GetRole(c *gin.Context) {
	roleID := c.Param("id")
	var role models.Role
	if err := configs.DB.First(&role, roleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch role")
		return
	}

	helpers.SuccessResponse(c, "Successfully fetched role", role)
}

func UpdateRole(c *gin.Context) {
	roleID := c.Param("id")
	var role models.Role
	if err := configs.DB.First(&role, roleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch role")
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := configs.DB.Save(&role).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to update role")
		return
	}

	helpers.SuccessResponse(c, "Role updated successfully", role)
}

func DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	var role models.Role
	if err := configs.DB.First(&role, roleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch role")
		return
	}

	if err := configs.DB.Delete(&role).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to delete role")
		return
	}

	helpers.SuccessResponse(c, "Role deleted successfully", nil)
}