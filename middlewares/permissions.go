package middlewares

import (
	"net/http"

	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/models"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permission string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.MustGet("user_id").(uint) 
        user := models.Users{}
        
        if err := configs.DB.Preload("Roles.Permissions").First(&user, userID).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
            c.Abort()
            return
        }

        hasPermission := false
        for _, role := range user.Roles {
            for _, perm := range role.Permissions {
                if perm.Name == permission {
                    hasPermission = true
                    break
                }
            }
            if hasPermission {
                break
            }
        }

        if !hasPermission {
            c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
            c.Abort()
            return
        }

        c.Next()
    }
}
