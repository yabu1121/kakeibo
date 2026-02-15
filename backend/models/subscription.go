package models

import (
	"time"
	"github.com/google/uuid"
)

type Subscription struct {
	BaseModel
	Name            string    `json:"name" gorm:"not null"`
	MonthlyFee      uint64    `json:"monthly_fee" gorm:"not null"`
	BilingCycleDays uint      `json:"biling_cycle_days" gorm:"not null"`
	NextBillingDate time.Time `json:"next_billing_date" gorm:"not null"`
	IsActive        bool      `json:"is_active" gorm:"not null"`

	UserID     uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	CategoryID uuid.UUID `json:"category_id" gorm:"type:char(36);not null;index"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
}