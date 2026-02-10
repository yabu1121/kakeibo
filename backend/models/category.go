package models

// Category カテゴリーテーブル
type Category struct {
	BaseModel
	Name string `json:"name" gorm:"not null;unique"`
}
