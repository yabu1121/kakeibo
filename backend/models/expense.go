package models

import (
	"time"
	"github.com/google/uuid"
)

type Expense struct {
	BaseModel
	Amount      int       `json:"amount" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	SpentAt     time.Time `json:"spent_at" gorm:"not null;index"`

	UserID     uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	CategoryID uuid.UUID `json:"category_id" gorm:"type:char(36);not null;index"`

	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
}
