package routes

import (
	"github.com/Fawaidac/siapel_backend/controllers"
	"github.com/Fawaidac/siapel_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
				auth.POST("/register", controllers.Register)
				auth.POST("/login", controllers.Login)
	}

	protected := router.Group("/user")
	protected.Use(middlewares.AuthMiddleware())
	{
				protected.GET("/profile", controllers.UserProfile)
				protected.POST("/logout", controllers.Logout)
	}

	subDistrict := router.Group("/sub-districts")
	subDistrict.Use(middlewares.AuthMiddleware())
	{
				subDistrict.POST("/", controllers.CreateSubDistrict)
				subDistrict.GET("/", controllers.GetAllSubDistricts)
				subDistrict.GET("/:id", controllers.GetSubDistrictByID)
				subDistrict.PUT("/:id", controllers.UpdateSubDistrict)
				subDistrict.DELETE("/:id", controllers.DeleteSubDistrict)
	}

	villageGroup := router.Group("/villages")
	subDistrict.Use(middlewares.AuthMiddleware())
	{
    villageGroup.POST("/", controllers.CreateVillage)
    villageGroup.GET("/", controllers.GetAllVillages)
    villageGroup.GET("/:id", controllers.GetVillageByID)
    villageGroup.GET("/select/:id_kecamatan", controllers.GetVillagesBySubDistrict)
    villageGroup.PUT("/:id", controllers.UpdateVillage)
    villageGroup.DELETE("/:id", controllers.DeleteVillage)
	}
	
}
