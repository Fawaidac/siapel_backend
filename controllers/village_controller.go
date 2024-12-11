package controllers

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/helpers"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
)

func CreateVillage(c *gin.Context) {
    var village models.Village
    if err := c.ShouldBindJSON(&village); err != nil {
        helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := configs.DB.Create(&village).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create village")
        return
    }

    helpers.SuccessResponse(c, "Village created successfully", village)
}

func GetAllVillages(c *gin.Context) {
    var villages []models.Village
    if err := configs.DB.Find(&villages).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve villages")
        return
    }

    helpers.SuccessResponse(c, "Villages retrieved successfully", villages)
}

func GetVillageByID(c *gin.Context) {
    id := c.Param("id")
    var village models.Village

    if err := configs.DB.First(&village, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusNotFound, "Village not found")
        return
    }

    helpers.SuccessResponse(c, "Village retrieved successfully", village)
}

func GetVillagesBySubDistrict(c *gin.Context) {
    idKecamatan := c.Param("id_kecamatan")
    var villages []models.Village

    if err := configs.DB.Where("id_kecamatan = ?", idKecamatan).Find(&villages).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve villages by sub-district")
        return
    }

    helpers.SuccessResponse(c, "Villages retrieved successfully by sub-district", villages)
}

func UpdateVillage(c *gin.Context) {
    id := c.Param("id")
    var village models.Village

    if err := configs.DB.First(&village, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusNotFound, "Village not found")
        return
    }

    if err := c.ShouldBindJSON(&village); err != nil {
        helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := configs.DB.Save(&village).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update village")
        return
    }

    helpers.SuccessResponse(c, "Village updated successfully", village)
}

func DeleteVillage(c *gin.Context) {
    id := c.Param("id")
    if err := configs.DB.Delete(&models.Village{}, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete village")
        return
    }

    helpers.SuccessResponse(c, "Village deleted successfully", nil)
}
