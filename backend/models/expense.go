package models

import (
	"time"
)

// Expense 消費テーブル
type Expense struct {
	BaseModel
	UserID     uint      `json:"user_id" gorm:"not null;index"`     // ユーザーID
	Amount     int       `json:"amount" gorm:"not null"`            // 値段
	Date       time.Time `json:"date" gorm:"not null;index"`        // 日付
	Memo       string    `json:"memo"`                              // メモ
	CategoryID uint      `json:"category_id" gorm:"not null;index"` // カテゴリーID

	// リレーション
	User     User     `json:"user" gorm:"foreignKey:UserID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}
