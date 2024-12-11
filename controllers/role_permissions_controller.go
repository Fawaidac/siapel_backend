package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AssignPermissionToRole(c *gin.Context) {
	var request struct {
		RoleID       uint `json:"role_id" binding:"required"`
		PermissionID uint `json:"permission_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var role models.Role
	var permission models.Permission
	if err := configs.DB.First(&role, request.RoleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := configs.DB.First(&permission, request.PermissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Permission not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := configs.DB.Create(&models.RolePermission{
		RoleID:       request.RoleID,
		PermissionID: request.PermissionID,
	}).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to assign permission")
		return
	}

	helpers.SuccessResponse(c, "Permission assigned to role successfully", nil)
}

func RevokePermissionFromRole(c *gin.Context) {
	var request struct {
		RoleID       uint `json:"role_id" binding:"required"`
		PermissionID uint `json:"permission_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var role models.Role
	var permission models.Permission
	if err := configs.DB.First(&role, request.RoleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := configs.DB.First(&permission, request.PermissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Permission not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := configs.DB.Where("role_id = ? AND permission_id = ?", request.RoleID, request.PermissionID).Delete(&models.RolePermission{}).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to revoke permission")
		return
	}

	helpers.SuccessResponse(c, "Permission revoked from role successfully", nil)
}

func AssignRoleToUser(c *gin.Context) {
	var request struct {
		UserID uint `json:"user_id" binding:"required"`
		RoleID uint `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.Users
	var role models.Role
	if err := configs.DB.First(&user, request.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "User not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := configs.DB.First(&role, request.RoleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := configs.DB.Create(&models.UserRole{
		UserID: request.UserID,
		RoleID: request.RoleID,
	}).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to assign role")
		return
	}

	helpers.SuccessResponse(c, "Role assigned to user successfully", nil)
}

func RevokeRoleFromUser(c *gin.Context) {
	var request struct {
		UserID uint `json:"user_id" binding:"required"`
		RoleID uint `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.Users
	var role models.Role
	if err := configs.DB.First(&user, request.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "User not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := configs.DB.First(&role, request.RoleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, http.StatusNotFound, "Role not found")
			return
		}
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := configs.DB.Where("user_id = ? AND role_id = ?", request.UserID, request.RoleID).Delete(&models.UserRole{}).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Unable to revoke role")
		return
	}

	helpers.SuccessResponse(c, "Role revoked from user successfully", nil)
}
