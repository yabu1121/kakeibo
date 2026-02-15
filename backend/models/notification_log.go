package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationLog struct {
	BaseModel
	SubscriptionRemind bool `json:"subscription_remind"`
	PublicFeeRemind bool `json:"public_fee_remind"`
	SentAt time.Time `json:"sent_at"`

	UserID uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
}