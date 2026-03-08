
package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TenantMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract role from JWT (Keycloak)
		// and find its associated database
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Mocked database lookup based on role (JWT integration needed)
		// In production, this would use the mapping table mentioned in AGENTS.md
		role := "TENANT_A" // Extracted from JWT
		tenantDBName := "go_duck_tenant_a" // From tenant_roles table

		// Store tenant info in context
		c.Set("tenantDB", tenantDBName)
		c.Next()
	}
}
