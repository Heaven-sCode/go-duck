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
Color string `json:"color" `
CreatedBy string `gorm:"column:created_by" json:"createdBy"`
CreatedDate time.Time `gorm:"column:created_date" json:"createdDate"`
LastModifiedBy string `gorm:"column:last_modified_by" json:"lastModifiedBy"`
LastModifiedDate time.Time `gorm:"column:last_modified_date" json:"lastModifiedDate"`
LastModifiedUserID string `gorm:"column:last_modified_user_id" json:"lastModifiedUserId"`
PersonID *uint `gorm:"column:person_id;index" json:"personId"`
Person *Person `gorm:"foreignKey:PersonID"
json:"person,omitempty"`
}