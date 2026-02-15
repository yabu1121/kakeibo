package models

import "github.com/google/uuid"

type Category struct {
	BaseModel
	Name string    `json:"name" gorm:"not null"`
}