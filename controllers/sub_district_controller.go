package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
)

func CreateSubDistrict(c *gin.Context) {
    var subDistrict models.SubDistrict
    if err := c.ShouldBindJSON(&subDistrict); err != nil {
        helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := configs.DB.Create(&subDistrict).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create sub-district")
        return
    }

    helpers.SuccessResponse(c, "Sub-district created successfully", subDistrict)
}

func GetAllSubDistricts(c *gin.Context) {
    var subDistricts []models.SubDistrict
    if err := configs.DB.Find(&subDistricts).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve sub-districts")
        return
    }

    helpers.SuccessResponse(c, "Sub-districts retrieved successfully", subDistricts)
}

func GetSubDistrictByID(c *gin.Context) {
    id := c.Param("id")
    var subDistrict models.SubDistrict

    if err := configs.DB.First(&subDistrict, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusNotFound, "Sub-district not found")
        return
    }

    helpers.SuccessResponse(c, "Sub-district retrieved successfully", subDistrict)
}

func UpdateSubDistrict(c *gin.Context) {
    id := c.Param("id")
    var subDistrict models.SubDistrict

    if err := configs.DB.First(&subDistrict, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusNotFound, "Sub-district not found")
        return
    }

    if err := c.ShouldBindJSON(&subDistrict); err != nil {
        helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := configs.DB.Save(&subDistrict).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update sub-district")
        return
    }

    helpers.SuccessResponse(c, "Sub-district updated successfully", subDistrict)
}

func DeleteSubDistrict(c *gin.Context) {
    id := c.Param("id")
    if err := configs.DB.Delete(&models.SubDistrict{}, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete sub-district")
        return
    }

    helpers.SuccessResponse(c, "Sub-district deleted successfully", nil)
}
