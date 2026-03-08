
package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TenantMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get roles from JWT (previously set by JWTMiddleware)
		userRolesInterface, exists := c.Get("UserRoles")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No roles found in token"})
			c.Abort()
			return
		}

		roles, ok := userRolesInterface.([]interface{})
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid roles format"})
			c.Abort()
			return
		}

		// 2. Lookup DB name for the roles in tenant_roles table
		var dbName string
		err := db.Raw("SELECT db_name FROM tenant_roles WHERE role_name IN ? LIMIT 1", roles).Scan(&dbName).Error
		
		if err != nil || dbName == "" {
			// Fallback to default DB if no role match (optional behavior)
			dbName = "go-duck" 
		}

		// 3. Store tenant info for downstream use
		c.Set("tenantDB", dbName)
		c.Next()
	}
}
