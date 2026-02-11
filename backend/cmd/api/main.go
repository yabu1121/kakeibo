package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// "gorm.io/gorm/logger" // This line was removed as per instruction

	"kakeibo-backend/handlers"
	"kakeibo-backend/models"
)

func main() {
	// データベース接続設定 (環境変数から取得)
	// Docker環境などでの接続待ちを考慮してリトライロジックを入れることが多いですが、
	// ここではシンプルに実装します。Docker Composeのdepends_onやwait-for-itを利用するのも手です。

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "kakeibo"),
		getEnv("DB_PORT", "5432"),
	)

	var db *gorm.DB
	var err error

	// 簡易的なリトライロジック
	for i := 0; i < 10; i++ {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v. Retrying in 2 seconds... (%d/10)", err, i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// マイグレーション実行
	// 実際の運用ではマイグレーションツール(golang-migrateなど)の使用を検討してください。
	if err := db.AutoMigrate(
		&models.User{},
		&models.Subscription{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Echoインスタンス作成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// ハンドラー初期化
	userHandler := handlers.UserHandlers{DB: db}
	subscriptionHandler := handlers.SubscriptionHandlers{DB: db}
	categoryHandler := handlers.CategoryHandlers{DB: db}
	expenseHandler := handlers.ExpenseHandlers{DB: db}

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo & Gorm!")
	})
	api := e.Group("/api")

	// User routes
	api.POST("/users", userHandler.CreateUser)
	api.GET("/users", userHandler.GetUsers)
	api.PUT("/users/:id", userHandler.UpdateUser)
	api.DELETE("/users/:id", userHandler.DeleteUser)

	// Subscription routes
	api.POST("/subscriptions", subscriptionHandler.CreateSubscription)
	api.GET("/subscriptions", subscriptionHandler.GetSubscription)
	api.PUT("/subscriptions/:id", subscriptionHandler.UpdateSubscription)
	api.DELETE("/subscriptions/:id", subscriptionHandler.DeleteSubscription)

	api.POST("/category", categoryHandler.CreateCategory)
	api.GET("/category", categoryHandler.GetCategory)
	api.PUT("/category", categoryHandler.UpdateCategory)
	api.DELETE("/category", categoryHandler.DeleteCategory)

	api.POST("/expense", expenseHandler.CreateExpense)
	api.GET("/expense", expenseHandler.GetExpense)
	api.PUT("/expense/:id", expenseHandler.UpdateExpense)
	api.DELETE("/expense/:id", expenseHandler.DeleteExpense)
	// サーバー起動
	port := getEnv("PORT", "8080")
	e.Logger.Fatal(e.Start(":" + port))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
