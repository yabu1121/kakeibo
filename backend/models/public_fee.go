package models

import (
	"time"

	"github.com/google/uuid"
)

type PublicFee struct {
	BaseModel
	FeeType string `json:"fee_type" gorm:"not null"`
	Amount  uint64 `json:"amount" gorm:"not null"`
	UsageMonth uint `json:"usage_month" gorm:"not null"`
	NextBillingDate time.Time `json:"next_billing_date" gorm:"not null"`

	UserID uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	CategoryID uuid.UUID `json:"category_id" gorm:"type:char(36);not null;index"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}