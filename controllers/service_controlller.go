package controllers

import (
    "net/http"

    "github.com/Fawaidac/siapel_backend/configs"
    "github.com/Fawaidac/siapel_backend/helpers"
    "github.com/Fawaidac/siapel_backend/models"
    "github.com/gin-gonic/gin"
)

// CreateService - Create a new service
func CreateService(c *gin.Context) {
    var service models.Service
    if err := c.ShouldBindJSON(&service); err != nil {
        helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := configs.DB.Create(&service).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create service")
        return
    }

    helpers.SuccessResponse(c, "Service created successfully", service)
}

// GetAllServices - Get all services
func GetAllServices(c *gin.Context) {
    var services []models.Service
    if err := configs.DB.Find(&services).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve services")
        return
    }

    helpers.SuccessResponse(c, "Services retrieved successfully", services)
}

// GetServiceByID - Get a single service by ID
func GetServiceByID(c *gin.Context) {
    id := c.Param("id")
    var service models.Service

    if err := configs.DB.First(&service, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusNotFound, "Service not found")
        return
    }

    helpers.SuccessResponse(c, "Service retrieved successfully", service)
}

// UpdateService - Update a service
func UpdateService(c *gin.Context) {
    id := c.Param("id")
    var service models.Service

    if err := configs.DB.First(&service, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusNotFound, "Service not found")
        return
    }

    if err := c.ShouldBindJSON(&service); err != nil {
        helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := configs.DB.Save(&service).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update service")
        return
    }

    helpers.SuccessResponse(c, "Service updated successfully", service)
}

// DeleteService - Delete a service
func DeleteService(c *gin.Context) {
    id := c.Param("id")
    if err := configs.DB.Delete(&models.Service{}, id).Error; err != nil {
        helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete service")
        return
    }

    helpers.SuccessResponse(c, "Service deleted successfully", nil)
}
