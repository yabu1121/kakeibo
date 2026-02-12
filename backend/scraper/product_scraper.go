package scraper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Product はスクレイピングで取得する商品情報
type Product struct {
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	ScrapedAt   time.Time `json:"scraped_at"`
}

// ScraperConfig はスクレイパーの設定
type ScraperConfig struct {
	AllowedDomains []string      // 許可するドメイン
	UserAgent      string        // User-Agent
	Delay          time.Duration // リクエスト間隔
	Parallelism    int           // 並行リクエスト数
	CacheDir       string        // キャッシュディレクトリ
}

// DefaultConfig はデフォルト設定を返す
func DefaultConfig() *ScraperConfig {
	return &ScraperConfig{
		AllowedDomains: []string{},
		UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		Delay:          2 * time.Second,
		Parallelism:    1,
		CacheDir:       "./cache",
	}
}

// ProductScraper は商品情報をスクレイピングする
type ProductScraper struct {
	collector *colly.Collector
	config    *ScraperConfig
	products  []Product
}

// NewProductScraper は新しいProductScraperを作成
func NewProductScraper(config *ScraperConfig) *ProductScraper {
	if config == nil {
		config = DefaultConfig()
	}

	c := colly.NewCollector(
		colly.UserAgent(config.UserAgent),
		colly.CacheDir(config.CacheDir),
		colly.Async(true),
	)

	// レート制限設定
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: config.Parallelism,
		Delay:       config.Delay,
	})

	if len(config.AllowedDomains) > 0 {
		c.AllowedDomains = config.AllowedDomains
	}

	return &ProductScraper{
		collector: c,
		config:    config,
		products:  []Product{},
	}
}

// ScrapeExample は例としてシンプルなスクレイピングを実行
// 実際のサイトに合わせてセレクタを変更してください
func (ps *ProductScraper) ScrapeExample(url string) ([]Product, error) {
	ps.products = []Product{}

	// リクエスト前のログ
	ps.collector.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL.String())
	})

	// エラーハンドリング
	ps.collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error visiting %s: %v", r.Request.URL, err)
	})

	// HTMLから商品情報を抽出
	// 注意: これは例です。実際のサイトに合わせてセレクタを変更してください
	ps.collector.OnHTML(".product-item", func(e *colly.HTMLElement) {
		product := Product{
			Name:        strings.TrimSpace(e.ChildText(".product-name")),
			Price:       strings.TrimSpace(e.ChildText(".product-price")),
			URL:         e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
			ImageURL:    e.ChildAttr("img", "src"),
			Description: strings.TrimSpace(e.ChildText(".product-description")),
			ScrapedAt:   time.Now(),
		}

		// 空のデータは追加しない
		if product.Name != "" {
			ps.products = append(ps.products, product)
			log.Printf("Found product: %s - %s", product.Name, product.Price)
		}
	})

	// 次のページへのリンクを辿る(ページネーション)
	ps.collector.OnHTML("a.next-page", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		if nextPage != "" {
			e.Request.Visit(nextPage)
		}
	})

	// スクレイピング完了時
	ps.collector.OnScraped(func(r *colly.Response) {
		log.Printf("Finished scraping: %s", r.Request.URL)
	})

	// スクレイピング開始
	err := ps.collector.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("failed to visit URL: %w", err)
	}

	// 非同期処理の完了を待つ
	ps.collector.Wait()

	return ps.products, nil
}

// ScrapeWithCustomSelector はカスタムセレクタでスクレイピング
func (ps *ProductScraper) ScrapeWithCustomSelector(
	url string,
	itemSelector string,
	nameSelector string,
	priceSelector string,
	linkSelector string,
) ([]Product, error) {
	ps.products = []Product{}

	ps.collector.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL.String())
	})

	ps.collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error: %v", err)
	})

	// カスタムセレクタで抽出
	ps.collector.OnHTML(itemSelector, func(e *colly.HTMLElement) {
		product := Product{
			Name:      strings.TrimSpace(e.ChildText(nameSelector)),
			Price:     strings.TrimSpace(e.ChildText(priceSelector)),
			URL:       e.Request.AbsoluteURL(e.ChildAttr(linkSelector, "href")),
			ScrapedAt: time.Now(),
		}

		if product.Name != "" {
			ps.products = append(ps.products, product)
		}
	})

	err := ps.collector.Visit(url)
	if err != nil {
		return nil, err
	}

	ps.collector.Wait()
	return ps.products, nil
}

// GetProducts は取得した商品リストを返す
func (ps *ProductScraper) GetProducts() []Product {
	return ps.products
}

// ClearProducts は商品リストをクリア
func (ps *ProductScraper) ClearProducts() {
	ps.products = []Product{}
}
