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
	users := []models.User{
		{
			Name:        "田中太郎",
			RealName:    "田中太郎",
			Email:       "tanaka@example.com",
			Password:    "password123",
			Icon:        "https://i.pravatar.cc/150?img=1",
			ProfileMemo: "家計簿アプリの管理者です。節約が趣味です。",
		},
		{
			Name:        "佐藤花子",
			RealName:    "佐藤花子",
			Email:       "sato@example.com",
			Password:    "password123",
			Icon:        "https://i.pravatar.cc/150?img=2",
			ProfileMemo: "投資と貯金に興味があります。",
		},
		{
			Name:        "鈴木一郎",
			RealName:    "鈴木一郎",
			Email:       "suzuki@example.com",
			Password:    "password123",
			Icon:        "https://i.pravatar.cc/150?img=3",
			ProfileMemo: "毎月の支出を見える化したいです。",
		},
		{
			Name:        "高橋美咲",
			RealName:    "高橋美咲",
			Email:       "takahashi@example.com",
			Password:    "password123",
			Icon:        "https://i.pravatar.cc/150?img=4",
			ProfileMemo: "家族の家計を管理しています。",
		},
		{
			Name:        "伊藤健太",
			RealName:    "伊藤健太",
			Email:       "ito@example.com",
			Password:    "password123",
			Icon:        "https://i.pravatar.cc/150?img=5",
			ProfileMemo: "副業の収支も管理したいです。",
		},
	}

	// ユーザーデータを挿入
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("ユーザー作成エラー (%s): %v", user.Email, err)
		} else {
			log.Printf("✓ ユーザー作成成功: %s (%s)", user.Name, user.Email)
		}
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
