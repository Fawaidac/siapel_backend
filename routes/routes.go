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
}
