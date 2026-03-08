package main

import (
"log"
"net/http"

"github.com/gin-gonic/gin"
"gorm.io/driver/postgres"
"gorm.io/gorm"
"go-duck/management"
"go-duck/middleware"
"go-duck/controllers"
)

func main() {
// Initialize master DB connection
dsn := "host=localhost user=postgres password=password dbname=go-duck port=5432 sslmode=disable"
masterDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatalf("Failed to connect to master database: %v", err)
}

r := gin.Default()

// Health Check
r.GET("/health", func(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{"status": "UP"})
})

// Management APIs (Run-time DB creation)
mgmt := r.Group("/management")
{
mgmt.POST("/db/create", management.CreateDatabaseAndMigrate(masterDB))
}

// Audit API (Global)
auditCtrl := controllers.AuditController{DB: masterDB}
r.GET("/audit", auditCtrl.GetLogs)

// Application APIs (with Multi-tenancy and Auditing)
api := r.Group("/api")
api.Use(middleware.TenantMiddleware(masterDB))
api.Use(middleware.AuditMiddleware(masterDB))
{
// Register entity controllers here
}

r.Run(":8080")
}