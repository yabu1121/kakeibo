package scraper

import (
	"testing"
	"time"
)

// TestDefaultConfig はデフォルト設定のテスト
func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.UserAgent == "" {
		t.Error("UserAgent should not be empty")
	}

	if config.Delay != 2*time.Second {
		t.Errorf("Expected delay 2s, got %v", config.Delay)
	}

	if config.Parallelism != 1 {
		t.Errorf("Expected parallelism 1, got %d", config.Parallelism)
	}
}

// TestNewProductScraper はスクレイパーの初期化テスト
func TestNewProductScraper(t *testing.T) {
	scraper := NewProductScraper(nil)

	if scraper == nil {
		t.Error("Scraper should not be nil")
	}

	if scraper.collector == nil {
		t.Error("Collector should not be nil")
	}

	if scraper.config == nil {
		t.Error("Config should not be nil")
	}
}

// TestClearProducts は商品リストのクリアテスト
func TestClearProducts(t *testing.T) {
	scraper := NewProductScraper(nil)

	// テストデータを追加
	scraper.products = []Product{
		{Name: "Test Product", Price: "1000円"},
	}

	if len(scraper.GetProducts()) != 1 {
		t.Error("Should have 1 product")
	}

	scraper.ClearProducts()

	if len(scraper.GetProducts()) != 0 {
		t.Error("Products should be empty after clear")
	}
}

// 実際のスクレイピングテストは、テスト用のHTMLサーバーを立てて行うのが望ましい
// ここでは構造のテストのみ実施
