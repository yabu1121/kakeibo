package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Report struct {
	BaseModel
	TargetMonth       time.Time      `json:"target_month" gorm:"not null;index"`
	TotalExpense      uint64         `json:"total_expense" gorm:"not null"`
	TotalPublicFee    uint64         `json:"total_public_fee" gorm:"not null"`
	TotalSubscription uint64         `json:"total_subscription" gorm:"not null"`
	CategoryBreakdown datatypes.JSON `json:"category_breakdown" gorm:"type:json"`

	UserID uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
}