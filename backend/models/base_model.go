package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel GORMモデルのベース（JSONタグ付き）
type BaseModel struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:char(36)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
