
package management

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DatabaseRequest struct {
	Role   string `json:"role" binding:"required"`
	DBName string `json:"db_name" binding:"required"`
}

func CreateDatabaseAndMigrate(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DatabaseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. CREATE DATABASE
		createDBSql := fmt.Sprintf("CREATE DATABASE %s", req.DBName)
		if err := db.Exec(createDBSql).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create database: " + err.Error()})
			return
		}

		// 2. Insert into roles mapping table
		insertRoleSql := "INSERT INTO tenant_roles (role_name, db_name) VALUES (?, ?)"
		if err := db.Exec(insertRoleSql, req.Role, req.DBName).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map role: " + err.Error()})
			return
		}

		// 3. Start Liquibase Migration
		// This executes the liquibase CLI to migrate the newly created database
		cmd := exec.Command("liquibase", "--url=jdbc:postgresql://localhost:5432/"+req.DBName, "--changeLogFile=changelog.xml", "update")
		if err := cmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start migration: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Database created and migration started for " + req.Role})
	}
}
