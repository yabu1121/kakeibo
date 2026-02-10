package models

// Subscription サブスクテーブル
type Subscription struct {
	BaseModel
	UserID      uint   `json:"user_id" gorm:"not null;index"`     // ユーザーID
	ProductName string `json:"product_name" gorm:"not null"`      // 商品名
	CategoryID  uint   `json:"category_id" gorm:"not null;index"` // カテゴリーID
	Amount      int    `json:"amount"`                            // 金額
	Frequency   string `json:"frequency" gorm:"not null"`         // 頻度 (例: "monthly", "yearly")

	// リレーション
	User     User     `json:"user" gorm:"foreignKey:UserID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}
