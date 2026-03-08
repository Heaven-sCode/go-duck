package models

import (
"time"
)

type Car struct {
ID uint `gorm:"primaryKey" json:"id"`
Name string `json:"name" binding:"required"`
Model string `json:"model" binding:"required"`
Year int `json:"year" binding:"required"`
Price float64 `json:"price" binding:"required"`
CreatedBy string `json:"createdBy"`
CreatedDate time.Time `json:"createdDate"`
LastModifiedBy string `json:"lastModifiedBy"`
LastModifiedDate time.Time `json:"lastModifiedDate"`
Person Person `json:"person"`
CreatedAt time.Time `json:"createdAt"`
UpdatedAt time.Time `json:"updatedAt"`
}