package models

import (
"time"
)

type Person struct {
ID uint `gorm:"primaryKey" json:"id"`
Name string `json:"name" binding:"required"`
Age int `json:"age" binding:"required"`
Email string `json:"email" binding:"required"`
CreatedBy string `gorm:"column:created_by" json:"createdBy"`
CreatedDate time.Time `gorm:"column:created_date" json:"createdDate"`
LastModifiedBy string `gorm:"column:last_modified_by" json:"lastModifiedBy"`
LastModifiedDate time.Time `gorm:"column:last_modified_date" json:"lastModifiedDate"`
LastModifiedUserID string `gorm:"column:last_modified_user_id" json:"lastModifiedUserId"`
Cars []Car `gorm:"foreignKey:PersonID" json:"cars"`
}