package models

import (
	"time"
)

// PublicUtility 公共料金テーブル
type PublicUtility struct {
	BaseModel
	UserID     uint      `json:"user_id" gorm:"not null;index"`     // ユーザーID
	Amount     int       `json:"amount" gorm:"not null"`            // 金額
	Date       time.Time `json:"date" gorm:"not null;index"`        // 日付
	CategoryID uint      `json:"category_id" gorm:"not null;index"` // カテゴリーID
	Memo       string    `json:"memo"`                              // メモ

	// リレーション
	User     User     `json:"user" gorm:"foreignKey:UserID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}
