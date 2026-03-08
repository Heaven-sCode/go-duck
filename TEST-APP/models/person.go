package models

import (
"time"
)

type Person struct {
ID uint `gorm:"primaryKey" json:"id"`
Name string `json:"name" binding:"required"`
Age int `json:"age" binding:"required"`
Email string `json:"email" binding:"required"`
CreatedBy string `json:"createdBy"`
CreatedDate time.Time `json:"createdDate"`
LastModifiedBy string `json:"lastModifiedBy"`
LastModifiedDate time.Time `json:"lastModifiedDate"`
Cars []Car `json:"cars"`
CreatedAt time.Time `json:"createdAt"`
UpdatedAt time.Time `json:"updatedAt"`
}