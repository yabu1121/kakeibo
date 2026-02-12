package models

import (
	"gorm.io/gorm"
)

// ScrapedItem はスクレイピングで取得したデータを保存するためのモデル
type ScrapedItem struct {
	gorm.Model
	URL         string `json:"url"`
	Name        string `json:"name"`
	Price       string `json:"price"` // サイトによって形式が異なるため文字列として保存
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	WebsiteName string `json:"website_name"` // どこのサイトから取得したか（ホスト名など）
}
