package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"kakeibo-backend/scraper"
)

// Example1: 基本的なスクレイピング
func example1_BasicScraping() {
	fmt.Println("=== Example 1: Basic Scraping ===")

	// スクレイパーを作成
	config := scraper.DefaultConfig()
	config.AllowedDomains = []string{"example.com"}
	config.Delay = 1 * time.Second

	ps := scraper.NewProductScraper(config)

	// スクレイピング実行
	// 注意: example.comは実際にはスクレイピングできません。実際のサイトに置き換えてください
	products, err := ps.ScrapeExample("https://example.com/products")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// 結果を表示
	fmt.Printf("Found %d products\n", len(products))
	for i, product := range products {
		fmt.Printf("%d. %s - %s\n", i+1, product.Name, product.Price)
		fmt.Printf("   URL: %s\n", product.URL)
	}
}

// Example2: カスタムセレクタでスクレイピング
func example2_CustomSelector() {
	fmt.Println("\n=== Example 2: Custom Selector ===")

	ps := scraper.NewProductScraper(nil)

	// カスタムセレクタを指定
	products, err := ps.ScrapeWithCustomSelector(
		"https://example.com/items",
		".item",         // 商品アイテムのセレクタ
		"h3.item-title", // 商品名のセレクタ
		".item-price",   // 価格のセレクタ
		"a.item-link",   // リンクのセレクタ
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// JSONで出力
	jsonData, _ := json.MarshalIndent(products, "", "  ")
	fmt.Println(string(jsonData))
}

// Example3: 実践的な使用例 - 価格監視
func example3_PriceMonitoring() {
	fmt.Println("\n=== Example 3: Price Monitoring ===")

	// 監視したい商品のURL
	targetURL := "https://example.com/product/12345"

	// 定期的にチェック(実際は別のgoroutineやcronで実行)
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	checkPrice := func() {
		ps := scraper.NewProductScraper(nil)
		products, err := ps.ScrapeExample(targetURL)
		if err != nil {
			log.Printf("Error checking price: %v", err)
			return
		}

		if len(products) > 0 {
			product := products[0]
			fmt.Printf("[%s] %s: %s\n",
				product.ScrapedAt.Format("2006-01-02 15:04:05"),
				product.Name,
				product.Price,
			)

			// 価格が閾値以下なら通知(実装例)
			// if isPriceLow(product.Price) {
			//     sendNotification(product)
			// }
		}
	}

	// 初回実行
	checkPrice()

	// 定期実行(この例では1回だけ)
	select {
	case <-ticker.C:
		checkPrice()
	case <-time.After(2 * time.Second):
		fmt.Println("Example finished")
	}
}

// Example4: 複数ページのスクレイピング
func example4_MultiplePages() {
	fmt.Println("\n=== Example 4: Multiple Pages ===")

	ps := scraper.NewProductScraper(nil)
	allProducts := []scraper.Product{}

	// ページ1〜3をスクレイピング
	for page := 1; page <= 3; page++ {
		url := fmt.Sprintf("https://example.com/products?page=%d", page)
		fmt.Printf("Scraping page %d...\n", page)

		products, err := ps.ScrapeExample(url)
		if err != nil {
			log.Printf("Error on page %d: %v", page, err)
			continue
		}

		allProducts = append(allProducts, products...)
		ps.ClearProducts() // 次のページ用にクリア

		// サーバーに負荷をかけないよう待機
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Total products found: %d\n", len(allProducts))
}

func main() {
	fmt.Println("Scraping Examples")
	fmt.Println("=================\n")

	// 注意: これらの例は実際には動作しません
	// 実際のサイトに合わせてURLとセレクタを変更してください

	// 例を実行(コメントアウトを外して使用)
	// example1_BasicScraping()
	// example2_CustomSelector()
	// example3_PriceMonitoring()
	// example4_MultiplePages()

	fmt.Println("\n実際に使用する際は:")
	fmt.Println("1. 対象サイトの利用規約を確認")
	fmt.Println("2. robots.txtを確認")
	fmt.Println("3. 適切なセレクタを設定")
	fmt.Println("4. レート制限を設定")
}
