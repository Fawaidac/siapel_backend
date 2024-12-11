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
				protected.GET("/", middlewares.PermissionMiddleware("read_user"), controllers.GetAllUsers)
				protected.POST("/logout", controllers.Logout)
	}

	subDistrict := router.Group("/sub-districts")
	subDistrict.Use(middlewares.AuthMiddleware())
	{
				subDistrict.POST("/",  middlewares.PermissionMiddleware("create_sub_district"), controllers.CreateSubDistrict)
				subDistrict.GET("/",  middlewares.PermissionMiddleware("read_sub_district"), controllers.GetAllSubDistricts)
				subDistrict.GET("/:id",  middlewares.PermissionMiddleware("read_sub_district"), controllers.GetSubDistrictByID)
				subDistrict.PUT("/:id",  middlewares.PermissionMiddleware("update_sub_district"), controllers.UpdateSubDistrict)
				subDistrict.DELETE("/:id", middlewares.PermissionMiddleware("delete_sub_district"), controllers.DeleteSubDistrict)
	}

	villageGroup := router.Group("/villages")
	villageGroup.Use(middlewares.AuthMiddleware())
	{
		villageGroup.POST("/", middlewares.PermissionMiddleware("create_village"), controllers.CreateVillage)
		villageGroup.GET("/", middlewares.PermissionMiddleware("read_village"), controllers.GetAllVillages)
		villageGroup.GET("/:id", middlewares.PermissionMiddleware("read_village"), controllers.GetVillageByID)
		villageGroup.GET("/select/:id_kecamatan", middlewares.PermissionMiddleware("read_village"), controllers.GetVillagesBySubDistrict)
		villageGroup.PUT("/:id", middlewares.PermissionMiddleware("update_village"), controllers.UpdateVillage)
		villageGroup.DELETE("/:id", middlewares.PermissionMiddleware("delete_village"), controllers.DeleteVillage)
	}

	serviceGroup := router.Group("/services")
	serviceGroup.Use(middlewares.AuthMiddleware())
	{
		serviceGroup.POST("/", middlewares.PermissionMiddleware("create_service"), controllers.CreateService)
		serviceGroup.GET("/", middlewares.PermissionMiddleware("read_services"), controllers.GetAllServices)
		serviceGroup.GET("/:id", middlewares.PermissionMiddleware("read_service"), controllers.GetServiceByID)
		serviceGroup.PUT("/:id", middlewares.PermissionMiddleware("edit_service"), controllers.UpdateService)
		serviceGroup.DELETE("/:id", middlewares.PermissionMiddleware("delete_service"), controllers.DeleteService)
	}
	
	roleGroup := router.Group("roles")
	roleGroup.Use(middlewares.AuthMiddleware())
	{
		roleGroup.POST("/", middlewares.PermissionMiddleware("create_role"), controllers.CreateRole)       
		roleGroup.GET("/", middlewares.PermissionMiddleware("read_role"), controllers.GetRoles)          
		roleGroup.GET("/:id", middlewares.PermissionMiddleware("read_role"), controllers.GetRole)        
		roleGroup.PUT("/:id", middlewares.PermissionMiddleware("update_role"), controllers.UpdateRole)     
		roleGroup.DELETE("/:id", middlewares.PermissionMiddleware("delete_role"), controllers.DeleteRole) 

		roleGroup.POST("/:role_id/assign-user", middlewares.PermissionMiddleware("assign_role_to_user"), controllers.AssignRoleToUser)
		roleGroup.POST("/:role_id/revoke-user", middlewares.PermissionMiddleware("revoke_role_from_user"), controllers.RevokeRoleFromUser)
	}

	permissionGroup := router.Group("permissions")
	permissionGroup.Use(middlewares.AuthMiddleware())
	{
		permissionGroup.POST("/", middlewares.PermissionMiddleware("create_permission"), controllers.CreatePermission)       
		permissionGroup.GET("/", middlewares.PermissionMiddleware("read_permission"), controllers.GetPermissions)          
		permissionGroup.GET("/:id", middlewares.PermissionMiddleware("read_permission"), controllers.GetPermission)        
		permissionGroup.PUT("/:id", middlewares.PermissionMiddleware("update_permission"), controllers.UpdatePermission)     
		permissionGroup.DELETE("/:id", middlewares.PermissionMiddleware("delete_permission"), controllers.DeletePermission) 

		permissionGroup.POST("/:permission_id/assign-role", middlewares.PermissionMiddleware("assign_permission_to_role"), controllers.AssignPermissionToRole)
		permissionGroup.POST("/:permission_id/revoke-role", middlewares.PermissionMiddleware("revoke_permission_from_role"), controllers.RevokePermissionFromRole)
	}

}
