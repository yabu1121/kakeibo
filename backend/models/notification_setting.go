package models

import "github.com/google/uuid"

type NotificationSetting struct {
	BaseModel
	EnableSubscription bool `json:"enable_subscription"` 
	EnablePublicFee bool `json:"enable_public_fee"` 
	RemindDayOfMonth int `json:"remind_day_of_month"` 

	UserID uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
}