
package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go-duck/models"
)

func AuditMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.Next()
			return
		}

		// Simplified auditing logic
		method := c.Request.Method
		path := c.Request.URL.Path
		
		// Map method to action
		action := "UPDATE"
		if method == http.MethodPost { action = "CREATE" }
		if method == http.MethodDelete { action = "DELETE" }

		// Mock user and IP
		userEmail := c.GetHeader("User-Email")
		if userEmail == "" { userEmail = "anonymous" }
		
		keycloakId := c.GetHeader("X-Keycloak-Id")
		clientIP := c.ClientIP()

		// Call next handlers
		c.Next()

		// Logic to capture entity ID and snapshot values would go here...
		// For now, track the action
		auditEntry := models.AuditLog{
			EntityName: path,
			Action:     action,
			ModifiedBy: userEmail,
			KeycloakID: keycloakId,
			ModifiedAt: time.Now(),
			ClientIP:   clientIP,
		}
		db.Create(&auditEntry)
	}
}
