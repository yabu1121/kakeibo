package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"kakeibo-backend/models"
)

func main() {
	// データベース接続設定
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "kakeibo"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}

	log.Println("サンプルデータの作成を開始します...")

	// 既存データをクリア（開発環境のみ）
	if err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
		log.Printf("警告: テーブルのクリアに失敗しました: %v", err)
	}

	// サンプルユーザーデータ
	// サンプルユーザーデータ
	var users []models.User
	for i := 1; i <= 10000; i++ {
		users = append(users, models.User{
			Name:        fmt.Sprintf("田中太郎%d", i),
			RealName:    "田中太郎",
			Email:       fmt.Sprintf("tanaka_%d@example.com", i),
			Password:    "password123",
			Icon:        fmt.Sprintf("https://i.pravatar.cc/150?img=%d", i),
			ProfileMemo: "家計簿アプリの管理者です。節約が趣味です。",
		})
	}

	// ユーザーデータを一括挿入（バッチサイズ100）
	if err := db.CreateInBatches(&users, 100).Error; err != nil {
		log.Printf("ユーザー作成エラー: %v", err)
	} else {
		log.Println("✓ ユーザー作成成功")
	}

	log.Println("サンプルデータの作成が完了しました！")
	log.Printf("作成されたユーザー数: %d", len(users))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
